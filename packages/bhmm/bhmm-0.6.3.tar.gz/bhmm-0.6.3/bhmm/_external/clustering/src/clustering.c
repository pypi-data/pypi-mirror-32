/* * Copyright (c) 2015, 2014 Computational Molecular Biology Group, Free University
 * Berlin, 14195 Berlin, Germany.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without modification,
 * are permitted provided that the following conditions are met:
 *
 *  * Redistributions of source code must retain the above copyright notice, this
 * list of conditions and the following disclaimer.
 *  * Redistributions in binary form must reproduce the above copyright notice,
 * this list of conditions and the following disclaimer in the documentation and/or
 * other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ``AS IS''
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
 * ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
 * ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

#define NO_IMPORT_ARRAY
#include <clustering.h>
#include <assert.h>

FLT euclidean_distance(FLT *SKP_restrict a, FLT *SKP_restrict b, size_t n, FLT *buffer_a, FLT *buffer_b)
{
    double sum;
    size_t i;

    sum = 0.0;
    for(i=0; i<n; ++i) {
        sum += (a[i]-b[i])*(a[i]-b[i]);
    }
    return sqrt(sum);
}

/*
FLT minRMSD_distance(FLT *SKP_restrict a, FLT *SKP_restrict b, size_t n, FLT *SKP_restrict buffer_a, FLT *SKP_restrict buffer_b)
{
    FLT msd;
    FLT trace_a, trace_b;

    memcpy(buffer_a, a, n*sizeof(FLT));
    memcpy(buffer_b, b, n*sizeof(FLT));

    inplace_center_and_trace_atom_major(buffer_a, &trace_a, 1, n/3);
    inplace_center_and_trace_atom_major(buffer_b, &trace_b, 1, n/3);
    msd = msd_atom_major(n/3, n/3, buffer_a, buffer_b, trace_a, trace_b, 0, NULL);
    return sqrt(msd);
}
*/

int c_assign(FLT *chunk, FLT *centers, npy_int32 *dtraj, char* metric, Py_ssize_t N_frames, Py_ssize_t N_centers, Py_ssize_t dim) {
    int ret;
    FLT d, mindist;
    size_t argmin;
    FLT *buffer_a, *buffer_b;
    FLT (*distance)(FLT*, FLT*, size_t, FLT*, FLT*);

    buffer_a = NULL; buffer_b = NULL;
    ret = ASSIGN_SUCCESS;

    /* init metric */
    if(strcmp(metric,"euclidean")==0) {
        distance = euclidean_distance;
    } /*else if(strcmp(metric,"minRMSD")==0) {
        distance = minRMSD_distance;
        buffer_a = malloc(dim*sizeof(FLT));
        buffer_b = malloc(dim*sizeof(FLT));
        if(!buffer_a || !buffer_b) {
            ret = ASSIGN_ERR_NO_MEMORY; goto error;
        }
    } */
    else {
        ret = ASSIGN_ERR_INVALID_METRIC;
        goto error;
    }

    /* do the assignment */
    {
        Py_ssize_t i,j;
//        #pragma omp for private(j, argmin, mindist)
        for(i = 0; i < N_frames; ++i) {
            #ifdef CLUSTERING_64
            mindist = DBL_MAX;
            #else
            mindist = FLT_MAX;
            #endif
            argmin = -1;
            for(j = 0; j < N_centers; ++j) {
                d = distance(&chunk[i*dim], &centers[j*dim], dim, buffer_a, buffer_b);
//				#pragma omp critical
            	{
                	if(d<mindist) { mindist = d; argmin = j; }
            	}
            }
            dtraj[i] = argmin;
        }
    }

error:
    free(buffer_a);
    free(buffer_b);
    return ret;
}

PyObject *assign(PyObject *self, PyObject *args) {

    PyObject *py_centers, *py_res;
    PyArrayObject *np_chunk, *np_centers, *np_dtraj;
    Py_ssize_t N_centers, N_frames, dim;
    FLT *chunk;
    FLT *centers;
    npy_int32 *dtraj;
    char *metric;

    py_centers = NULL; py_res = NULL;
    np_chunk = NULL; np_dtraj = NULL;
    centers = NULL; metric=""; chunk = NULL; dtraj = NULL;

    if (!PyArg_ParseTuple(args, "O!OO!s", &PyArray_Type, &np_chunk, &py_centers, &PyArray_Type, &np_dtraj, &metric)) goto error; /* ref:borr. */

    /* import chunk */
    #ifdef CLUSTERING_64
    if(PyArray_TYPE(np_chunk)!=NPY_FLOAT64) { PyErr_SetString(PyExc_ValueError, "dtype of \"chunk\" isn\'t FLT (32)."); goto error; };
    #else
    if(PyArray_TYPE(np_chunk)!=NPY_FLOAT32) { PyErr_SetString(PyExc_ValueError, "dtype of \"chunk\" isn\'t FLT (32)."); goto error; };
    #endif
    if(!PyArray_ISCARRAY_RO(np_chunk) ) { PyErr_SetString(PyExc_ValueError, "\"chunk\" isn\'t C-style contiguous or isn\'t behaved."); goto error; };
    if(PyArray_NDIM(np_chunk)!=2) { PyErr_SetString(PyExc_ValueError, "Number of dimensions of \"chunk\" isn\'t 2."); goto error;  };
    N_frames = np_chunk->dimensions[0];
    dim = np_chunk->dimensions[1];
    if(dim==0) {
        PyErr_SetString(PyExc_ValueError, "chunk dimension must be larger than zero.");
        goto error;
    }
    chunk = PyArray_DATA(np_chunk);

    /* import dtraj */
    if(PyArray_TYPE(np_dtraj)!=NPY_INT32) { PyErr_SetString(PyExc_ValueError, "dtype of \"dtraj\" isn\'t int (32)."); goto error; };
    if(!PyArray_ISBEHAVED_RO(np_dtraj) ) { PyErr_SetString(PyExc_ValueError, "\"dtraj\" isn\'t behaved."); goto error; };
    if(PyArray_NDIM(np_dtraj)!=1) { PyErr_SetString(PyExc_ValueError, "Number of dimensions of \"dtraj\" isn\'t 1."); goto error; };
    if(np_chunk->dimensions[0]!=N_frames) {
        PyErr_SetString(PyExc_ValueError, "Size of \"dtraj\" differs from number of frames in \"chunk\".");
        goto error;
    }
    dtraj = (npy_int32*)PyArray_DATA(np_dtraj);

    /* import list of cluster centers */
    #ifdef CLUSTERING_64
    np_centers = (PyArrayObject*)PyArray_ContiguousFromAny(py_centers, NPY_FLOAT64, 2, 2);
    #else
    np_centers = (PyArrayObject*)PyArray_ContiguousFromAny(py_centers, NPY_FLOAT32, 2, 2);
    #endif
    if(!np_centers) {
        PyErr_SetString(PyExc_ValueError, "Could not convert \"centers\" to two-dimensional C-contiguous behaved ndarray of FLT (32).");
        goto error;
    }
    N_centers = np_centers->dimensions[0];
    if(N_centers==0) {
        PyErr_SetString(PyExc_ValueError, "centers must contain at least one element.");
        goto error;
    }
    if(np_centers->dimensions[1]!=dim) {
        PyErr_SetString(PyExc_ValueError, "Dimension of cluster centers doesn\'t match dimension of frames.");
        goto error;
    }
    centers = (FLT*)PyArray_DATA(np_centers);

    /* do the assignment */
    switch(c_assign(chunk, centers, dtraj, metric, N_frames, N_centers, dim)) {
        case ASSIGN_ERR_INVALID_METRIC:
            PyErr_SetString(PyExc_ValueError, "metric must be one of \"euclidean\" or \"minRMSD\".");
            goto error;
        case ASSIGN_ERR_NO_MEMORY:
            PyErr_NoMemory();
            goto error;
    }

    py_res = Py_BuildValue(""); /* =None */
    /* fall through */
error:
    return py_res;
}
