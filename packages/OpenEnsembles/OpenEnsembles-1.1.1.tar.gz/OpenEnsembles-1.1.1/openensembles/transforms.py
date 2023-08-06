"""
OpenEnsembles is a resource for performing and analyzing ensemble clustering
This file contains all transform functions. Each transform takes a data matrix
an x-vector and variable arguments. It also returns a data matrix, an x-vector
and a dictionary of parameters used (name, value). 

Copyright (C) 2017 Naegle Lab

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

"""

import numpy as np 
import pandas as pd 
import sklearn.cluster as skc
import matplotlib.pyplot as plt
from sklearn import datasets
import scipy.cluster.hierarchy as sch
from sklearn import preprocessing
import scipy.stats as stats
from types import FunctionType
import collections
import re
from sklearn.decomposition import PCA
import random

class transforms:
	"""
	Transform the data matrix according to the transformation procedure used and the optional arguments passed

	Parameters
	----------
	x: list
		The x-vector
	data: matrix
		The data matrix

	Attributes
	----------
	x_out: list
		The output x-vector, same as input x, unless transformation alters meaning of independent variables (such as PCA)
	data_out: matrix
		The transformed data matrix
	var_params: dict
		Listing of parameter used to create the transformation

	Other Parameters
	----------------
	These **kwargs depend on the specific transformation used. 

	See Also
	--------
	openensembles.data.transform()

	"""
	def __init__(self, x, data, kwargs):

		self.x = x
		self.data = data
		self.args = kwargs
		self.x_out = []
		self.data_out = []
		self.var_params = {}

	def transforms_available(self):
		""" 
		Get available transformations 

		Returns
		-------
		methods: dict
			Available methods returned as keys in dict
		"""
		methods =  [method for method in dir(self) if isinstance(getattr(self, method), collections.Callable)]
		methods.remove('transforms_available')
		methodDict = {}
		for method in methods:
			if not re.match('__', method):
				methodDict[method] = ''
		return methodDict


	def zscore(self):
		"""
		Uses stats.zscore to zscore along the axis given. By default, OpenEnsembles assumes that the
		feature dimensions are on axis=0 (column entries). Updates x_out accordingly

		Other Parameters
		----------------
		axis: {'both', 0, 1} (default=0)
			axis to operate on (default operates along column entries)
		
		Raises
		------
		ValueError:
			axis is not of type allowed
			   
		"""

		if 'axis' in self.args:
			axis = self.args['axis']
		else: 
			axis = 0

		self.x_out = self.x 
		
		if axis=='both':
			self.data_out = stats.zscore(self.data)
		elif axis==0 or axis==1:
			self.data_out = stats.zscore(self.data, axis)
		else:
			raise ValueError( "zscore must operate on columns (axis=0), rows (axis=1), or both (axis=''), you passed%s"%(axis))
		self.var_params = {}


	def minmax(self):
		"""
		Uses MinMaxScaler from `sklearn.preprocessing <http://scikit-learn.org/stable/modules/generated/sklearn.preprocessing.MinMaxScaler.html>`_ to scale data between a minimum value and 
		a maximum value. Defaults range from 0 to 1. Updates data_out accordingly.

		Other Parameters
		----------------
		minValue: float (default=0)
		maxValue: float (default=1)

		Raises
		------
		ValueError:
			If minValue > maxValue



		"""
		if 'minValue' in self.args:
			minValue = self.args['minValue']
		else:
			minValue = 0
		if 'maxValue' in self.args:
			maxValue = self.args['maxValue']
		else:
			maxValue = 1

		if minValue > maxValue:
			raise ValueError("Your requested minValue (%0.2f) is larger than the maximum value (%0.2f)"%(minValue, maxValue))

		self.x_out = self.x
		min_max_scaler = preprocessing.MinMaxScaler(feature_range=(minValue, maxValue))
		self.data_out = np.transpose(min_max_scaler.fit_transform(np.transpose(self.data)))
		self.var_params = {'minValue':minValue, 'maxValue':maxValue}
		
	def log(self):
		"""
		Log transformation will default to taking the log2 of all elements in the matrix. 

		Warnings
		--------
		Check for the creation of infinite values

		Parameters
		----------
		base: {2, 10, 'e', 'ln'} (default=2)

		Raises
		------
		ValueError: 
			if base type is not recognized

		"""
		if 'base' in self.args:
			base = self.args['base']
		else:
			base = 2 
		self.var_params['base'] = base
		self.x_out = self.x
		if isinstance(base, str):
			if base == 'e' or base=='ln':
				self.data_out = np.log(self.data)
		elif int(base) == 10:
			self.data_out = np.log10(self.data)
		elif int(base) == 2:
			self.data_out = np.log2(self.data)
		else:
			raise ValueError('Requested base for logarithm was not recognized as either e, 2, or 10)')


	def PCA(self):
		"""
		Applies `sklearn's decomposition by principal components (PCA) 
		<http://scikit-learn.org/stable/modules/generated/sklearn.decomposition.PCA.html>`_
		
		Applies PCA to data matrix. If variable argument n_components is set, it will keep the first n_components of the 
		post-transformed data.
		set n_components to a number between 0 and 1 to reduce dimensionality based on %variance explained.
		Updates data_out accordingly as the scores (or transformation of data objects into PCA space). 
		This also appends explained_variance to an attribute of class
		

		Other Parameters
		----------------
		n_components: int
			Number of components to keep, defaults to length of original feature vector

		Returns
		-------
		pca: PCA data object
			The intact pca object, such that one could retreive other parts of this

		
		"""

		if 'n_components' in self.args:
			n_components = self.args['n_components']
		else:
			n_components = self.data.shape[1]
		self.var_params['n_components'] = n_components
		pca = PCA(n_components=n_components)
		pca.fit(self.data)
		self.data_out = pca.transform(self.data)
		self.x_out = []
		for i in range(0, self.data_out.shape[1]):
			self.x_out.append("PC%d"%(i+1))
		#extra parameters
		self.explained_variance = pca.explained_variance_
			
		return pca

	def boxcox(self):
		"""
		This applies the boxcox transform to the data 


		Other Parameters
		----------------
		lambda: {None, scalar}
			The lambda value to use in the exponential, if 0, this is equivalent to taking the log. If None, 
			this fits the best lambda and returns the value. Default is None

		alpha: {None, float}, optional
			If alpha is not None, return the 100 * (1-alpha)% confidence interval for lmbda. Must be between 0.0 and 1.0.
		
		axis: {0, 1}
			Axis to boxcox on. 0 for rows, 1 for columns. Default is rows

		Returns
		-------
		lambdas: list of scalars
			The lambda values that maximized the log-likelihood function, one entry for each vector that was normalized

		"""

		if 'axis' not in self.args:
			axis = 0
		else:
			axis = self.args['axis']

		if 'lambda' not in self.args:
			lbda = None
		else:
			lbda = self.args['lambda']


		if 'alpha' not in self.args:
			alpha = None
		else:
			alpha = self.args['alpha']
			
		shape = self.data.shape

		D = np.zeros((shape[0], shape[1]))
		lambdas = []

		if axis == 0:
			for i in range(0, shape[0]):
				vec = self.data[i,:]
				vecOut, l = stats.boxcox(vec, lmbda = lbda, alpha=alpha)
				D[i,:] = vecOut
				lambdas.append(l)
		else:
			for i in range(0, shape[1]):
				vec = self.data[:,i]
				vecOut, l = stats.boxcox(vec, lmbda = lbda, alpha=alpha)
				D[:,i] = vecOut
				lambdas.append(l)
		self.data_out = D
		return(lambdas)


	def random_subsample(self):
		"""
		This returns a new data matrix that is a random subsampling of available features (columns)
		If it is the same length as original, then this will randomly reshuffle dimensions if sorted=True
		
		Other Parameters
		-----------------
		num_to_sample: int
			Number of random features to keep in output matrix 

		sorted: {True, False} default=True
			Whether to sort the random indexes in ascending order (keeping semblence of order of original features)

		"""
		if 'sorted' not in self.args:
			sorted=True
		else:
			sorted = self.args['sorted']

		if 'num_to_sample' not in self.args:
			raise ValueError("Expect a num_to_sample of new random dimensions")
		else:
			num_to_sample = self.args['num_to_sample']

		if num_to_sample < 1: 
			raise ValueError("Must keep at least one dimension in subsampling of features")

		if num_to_sample > len(self.x):
			raise ValueError("Cannot subsample with same or more dimensions, requested %d from %d available "%(num_to_sample, len(self.x)))


		#randomly select indexes to keep for x and data matrix
		indexesToKeep = random.sample(range(0, len(self.x)), num_to_sample)
		if sorted:
			indexesToKeep.sort()

		self.data_out = self.data[:,indexesToKeep]

		self.x_out = []
		for i in range (0, len(indexesToKeep)):
			self.x_out.append(self.x[indexesToKeep[i]]) 
		self.params = self.args


	def add_offset(self):
		"""
		This adds an offset (positive or negative) to all values in the matrix. This can be used
		to add floating point noise (i.e. create nono-zero values for log transform) or subtract the mean
		without requiring standard deviation scaling as in the zscore.
		
		Other Parameters
		-----------------
		offset: list of floats
			The value to add (use negative value for subtraction). If a single float is passed, it removes 
			the same value from the entire matrix. Otherwise, a list of values the same size as the feature 
			dimensions or the number of objects should be passed. 

		axis: int {0 for rows, 1 for columns}
			The dimension on which to subtract if offset is a list

		"""
		if 'offset' not in self.args:
			raise ValueError("Expect an offset value")
		else: 
			offset = self.args['offset']

		shape = self.data.shape
		if 'axis' in self.args:
			dimension = self.args['axis']
			if dimension != 0 and dimension != 1:
				raise ValueError('Dimension to operate in must be 0 (rows) or 1 (columns), you passed %d'%(dimension))
			if shape[int(not dimension)] != len(offset):
				raise ValueError("Length of offset, %d, is not the same as dimensionality %d"%(len(offset), shape[int(not dimension)]))

		if len(offset) == 1:
			self.data_out = self.data + offset[0]
		else:
			if 'axis' not in self.args:
				raise ValueError('argument axis required, 0 for row, 1 for column')
			#replicate the offset vector in either dimension and do matrix addition
			if dimension == 0: #operate on rows (meaning the offset better be length of columns)
				offset_array = np.tile(offset, [shape[0], 1])
				self.data_out = self.data + offset_array

			elif dimension == 1:
				array = np.asmatrix(offset)
				offset_array = np.tile(array.transpose(),[1,shape[1]])
				self.data_out = self.data + offset_array






	def internal_normalization(self):
		"""
		This normalizes all data (in rows) to one point in that row, based on either
		col_index or value in the x vector (x_val) (if both are passed, uses col_index)

		Other Parameters
		----------------
		col_index: int
			index in feature vector to use for normalization 
		x_val: float
			find and use the index of name x_val in the x vector. Recall that x was converted to ints if original features were string labels.



		Example:
			Normalize data to 5minutes internal_normalization(x_val=5)
		"""
		if 'col_index' in self.args:
			index = self.args['col_index']
			if index >= len(self.x):
				raise ValueError('internal_normalization requires an index value within the length of the x vector')
			x_val = self.x[index]

		elif 'x_val' in self.args:
			#find col_name
			x_val = self.args['x_val']
			indexes = np.flatnonzero(np.asarray(self.x) == x_val)
			if not indexes:
				raise ValueError('internal_normalization requires an x_val in the x vector, %s was not found'%(x_val))
			elif len(indexes) > 1:
				raise ValueError('internal_normalization requires an x_val in the x vector that appears one time, %s was found %d tiems'%(x_val, len(indexes)))
			else:
				index = indexes[0]
		else:
			raise ValueError('internal_normalization requires a valid normalizing column according to a specific value in x (x_val) or index (col_index) ')

		self.var_params['x_val'] = x_val
		self.var_params['index'] = index
		self.x_out = self.x
		vec = self.data[:,index]
		self.data_out = self.data/vec[:,None]




		

