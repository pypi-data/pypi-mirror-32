"""
proximity.py
---------------

Query mesh- point proximity.
"""
import numpy as np

from . import util

from .constants import tol, _log_time
from .triangles import closest_point as closest_point_corresponding

from collections import deque


def nearby_faces(mesh, points):
    """
    For each point find nearby faces relativly quickly.

    The closest point on the mesh to the queried point is guaranteed to be
    on one of the faces listed.

    Does this by finding the nearest vertex on the mesh to each point, and
    then returns all the faces that intersect the axis aligned bounding box
    centered at the queried point and extending to the nearest vertex.

    Parameters
    ----------
    mesh : Trimesh object
    points : (n,3) float , points in space

    Returns
    -----------
    candidates : (points,) int, sequence of indexes for mesh.faces
    """
    points = np.asanyarray(points, dtype=np.float64)
    if not util.is_shape(points, (-1, 3)):
        raise ValueError('points must be (n,3)!')

    # an r-tree containing the axis aligned bounding box for every triangle
    rtree = mesh.triangles_tree
    # a kd-tree containing every vertex of the mesh
    kdtree = mesh.kdtree

    # query the distance to the nearest vertex to get AABB of a sphere
    distance_vertex = kdtree.query(points)[0].reshape((-1, 1))
    distance_vertex += tol.merge

    # axis aligned bounds
    bounds = np.column_stack((points - distance_vertex,
                              points + distance_vertex))

    # faces that intersect axis aligned bounding box
    candidates = [list(rtree.intersection(b)) for b in bounds]

    return candidates


def closest_point_naive(mesh, points):
    """
    Given a mesh and a list of points, find the closest point on any triangle.

    Does this by constructing a very large intermediate array and
    comparing every point to every triangle.

    Parameters
    ----------
    triangles : (n,3,3) float, triangles in space
    points    : (m,3)   float, points in space

    Returns
    ----------
    closest     : (m,3) float, closest point on triangles for each point
    distance    : (m,)  float, distance
    triangle_id : (m,)  int, index of triangle containing closest point
    """

    # establish that input triangles and points are sane
    triangles = mesh.triangles.view(np.ndarray)
    points = np.asanyarray(points, dtype=np.float64)
    if not util.is_shape(triangles, (-1, 3, 3)):
        raise ValueError('triangles shape incorrect')
    if not util.is_shape(points, (-1, 3)):
        raise ValueError('points must be (n,3)')

    # create a giant tiled array of each point tiled len(triangles) times
    points_tiled = np.tile(points, (1, len(triangles)))
    on_triangle = np.array([closest_point_corresponding(
        triangles, i.reshape((-1, 3))) for i in points_tiled])

    # distance squared
    distance_2 = [((i - q)**2).sum(axis=1)
                  for i, q in zip(on_triangle, points)]
    triangle_id = np.array([i.argmin() for i in distance_2])

    # closest cartesian point
    closest = np.array([g[i] for i, g in zip(triangle_id, on_triangle)])
    distance = np.array([g[i] for i, g in zip(triangle_id, distance_2)]) ** .5

    return closest, distance, triangle_id


def closest_point(mesh, points):
    """
    Given a mesh and a list of points, find the closest point on any triangle.

    Parameters
    ----------
    mesh   : Trimesh object
    points : (m,3)   float, points in space

    Returns
    ----------
    closest     : (m,3) float, closest point on triangles for each point
    distance    : (m,)  float, distance
    triangle_id : (m,)  int, index of triangle containing closest point
    """

    points = np.asanyarray(points, dtype=np.float64)
    if not util.is_shape(points, (-1, 3)):
        raise ValueError('points must be (n,3)!')

    # do a tree- based query for faces near each point
    candidates = nearby_faces(mesh, points)
    # view triangles as an ndarray so we don't have to recompute
    # the MD5 during all of the subsequent advanced indexing
    triangles = mesh.triangles.view(np.ndarray)

    # create the corresponding list of triangles
    # and query points to send to the closest_point function
    query_point = deque()
    query_tri = deque()
    for triangle_ids, point in zip(candidates, points):
        query_point.append(np.tile(point, (len(triangle_ids), 1)))
        query_tri.append(triangles[triangle_ids])

    # stack points into an (n,3) array
    query_point = np.vstack(query_point)
    # stack triangles into an (n,3,3) array
    query_tri = np.vstack(query_tri)

    # do the computation for closest point
    query_close = closest_point_corresponding(query_tri, query_point)
    query_group = np.cumsum(np.array([len(i) for i in candidates]))[:-1]

    distance_2 = ((query_close - query_point) ** 2).sum(axis=1)

    # find the single closest point for each group of candidates
    result_close = np.zeros((len(points), 3), dtype=np.float64)
    result_tid = np.zeros(len(points), dtype=np.int64)
    result_distance = np.zeros(len(points), dtype=np.float64)

    # go through results to get minimum distance result
    for i, close_points, distance, candidate in zip(
            np.arange(len(points)),
            np.array_split(query_close, query_group),
            np.array_split(distance_2, query_group),
            candidates):

        # unless some other check is true use the smallest distance
        idx = distance.argmin()

        # if we have multiple candidates check them
        if len(candidate) > 1:
            # (2, ) int, list of 2 closest candidate indices
            idxs = distance.argsort()[:2]
            # make sure the two distances are identical
            check_distance = distance[idxs].ptp() < tol.merge
            # make sure the magnitude of both distances are nonzero
            check_magnitude = (np.abs(distance[idxs]) > tol.merge).all()

            # check if query-points are actually off-surface
            if check_distance and check_magnitude:
                # get face normals for two points
                normals = mesh.face_normals[np.array(candidate)[idxs]]
                # compute normalized surface-point to query-point vectors
                vectors = ((points[i] - close_points[idxs]) /
                           distance[idxs, np.newaxis] ** 0.5)
                # compare enclosed angle for both face normals
                dots = util.diagonal_dot(normals, vectors)
                # take the idx with the most positive angle
                idx = idxs[dots.argmax()]

        # take the single closest value from the group of values
        result_close[i] = close_points[idx]
        result_tid[i] = candidate[idx]
        result_distance[i] = distance[idx]

    # we were comparing the distance squared so
    # now take the square root in one vectorized operation
    result_distance **= .5

    return result_close, result_distance, result_tid


def signed_distance(mesh, points):
    """
    Find the signed distance from a mesh to a list of points.

    * Points OUTSIDE the mesh will have NEGATIVE distance
    * Points within tol.merge of the surface will have POSITIVE distance
    * Points INSIDE the mesh will have POSITIVE distance

    Parameters
    -----------
    mesh   : Trimesh object
    points : (n,3) float, list of points in space

    Returns
    ----------
    signed_distance : (n,3) float, signed distance from point to mesh
    """
    # make sure we have a numpy array
    points = np.asanyarray(points, dtype=np.float64)

    # find the closest point on the mesh to the queried points
    closest, distance, triangle_id = closest_point(mesh, points)

    # we only care about nonzero distances
    nonzero = distance > tol.merge

    if not nonzero.any():
        return distance

    inside = mesh.ray.contains_points(points[nonzero])
    sign = (inside.astype(int) * 2) - 1

    # apply sign to previously computed distance
    distance[nonzero] *= sign

    return distance


class ProximityQuery(object):
    """
    Proximity queries for the current mesh.
    """

    def __init__(self, mesh):
        self._mesh = mesh

    @_log_time
    def on_surface(self, points):
        """
        Given list of points, for each point find the closest point
        on any triangle of the mesh.

        Parameters
        ----------
        points : (m,3) float, points in space

        Returns
        ----------
        closest     : (m,3) float, closest point on triangles for each point
        distance    : (m,)  float, distance
        triangle_id : (m,)  int, index of closest triangle for each point
        """
        return closest_point(mesh=self._mesh,
                             points=points)

    def vertex(self, points):
        """
        Given a set of points, return the closest vertex index to each point

        Parameters
        ----------
        points : (n,3) float, list of points in space

        Returns
        ----------
        distance  : (n,) float, distance from source point to vertex
        vertex_id : (n,) int, index of mesh.vertices which is closest
        """
        tree = self._mesh.kdtree
        return tree.query(points)

    def signed_distance(self, points):
        """
        Find the signed distance from a mesh to a list of points.

        * Points OUTSIDE the mesh will have NEGATIVE distance
        * Points within tol.merge of the surface will have POSITIVE distance
        * Points INSIDE the mesh will have POSITIVE distance

        Parameters
        -----------
        points : (n,3) float, list of points in space

        Returns
        ----------
        signed_distance : (n,3) float, signed distance from point to mesh
        """
        return signed_distance(self._mesh, points)
