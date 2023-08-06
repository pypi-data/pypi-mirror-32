# Grimlock

We all know that when it comes to machine learning, it takes far more time
to preprocess your data than it does to actually build a model. Enter, <strong>grimlock</strong>.

<strong>grimlock</strong> will fix your missing values, handle data encoding, and feature scaling.

## Installation
Provided you already have NumPy, SciPy, Sci-kit Learn and Pandas already installed, the `grimlock` package is `pip`-installable:

<pre>$ pip install grimlock</pre>


### Cleaning Missing Data
Mesh of pandas.fillna() and sklearn Imputer

<pre>from grimlock import clean_missing
clean_missing(dataframe, column, clean_type='zero')</pre>


#### Parameters
- dataframe: dataframe variable
- column: column name (string)
- clean_type: 'zero' (default), 'mean', 'mode', 'most_frequent' (string)


### Convert Categorical
Quick conversion for categorical features (non-ordinal) 

<pre>from grimlock import convert_categorical
convert_categorical(dataframe, column, target_column)</pre>

#### Parameters
- dataframe: dataframe variable
- column: column name (string)
- target_column: target column name (string)



### Feature Scaling
<em>coming soon</em>