
# emtrees
Tree-based machine learning classifiers for embedded systems.
Train in Python, deploy on microcontroller.

Want Naive Bayes instead? Go to [embayes](https://github.com/jonnor/embayes)

## Key features

Embedded-friendly Classifier

* Portable C99 code
* No stdlib required
* No dynamic allocations
* Integer/fixed-point math only
* Single header file include
* Fast, sub-millisecond classification

Convenient Training

* API-compatible with [scikit-learn](http://scikit-learn.org)
* Implemented in Python 3
* C classifier accessible in Python using pybind11

[MIT licensed](./LICENSE.md)

## Status
**Minimally useful**

* Random Forests and ExtraTrees classifiers implemented
* Tested running on AVR, ESP8266 and Linux.
* On ESP8266, 8x8 digits classify in under 0.3ms with 95%+ accuracy
* On Linux, is approx 2x faster than sklearn

## Installing

Install from PyPI

    pip install emtrees --user

## Usage

1. Train your model in Python

```python
import emtrees
estimator = emtrees.RandomForest(n_estimators=10, max_depth=10)
estimator.fit(X_train, Y_train)
...
```

2. Generate C code
```python
code = estimator.output_c('sonar')
with open('sonar.h', 'w') as f:
   f.write(code)
```

3. Use the C code

```c
#include <emtrees.h>
#include "sonar.h"

const int32_t length = 60;
int32_t values[length] = { ... };
const int32_t predicted_class = sonar_predict(values, length):
```

For full example code, see [examples/digits.py](./examples/digits.py)
and [emtrees.ino](./emtrees.ino)

## TODO

0.2

* Standalone example application on microcontroller

1.0

* Support returning probabilities
* Support serializing/deserializing trees

Maybe

* Support weighted voting
* Support GradientBoostingClassifier
* Implement a Very Fast Decision Tree (VFDT) learning algorithm
* Support [XGBoost](https://github.com/dmlc/xgboost) learning of trees
* Support [LightGBM](https://github.com/Microsoft/LightGBM) learning of trees
* Implement multithreading when used in Python bindings, using OpenMP
* Support regression trees
