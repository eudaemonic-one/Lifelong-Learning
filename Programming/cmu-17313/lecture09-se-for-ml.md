# Lecture 9 Software Engineering for ML/AI

## ML Development

* Observation
* Hypothesis
* Predict
* Test
* Reject or Refine Hypothesis

### Three Fundamental Differences

* Data discovery and management
* Customization and Reuse
* No incremental development of model itself

## Machine Learning Pipeline

* Static
  * Get labeled data (data collection, cleaning, and labeling)
  * Identify and extract features (feature engineering)
  * Split data into training and evaluation set
  * Learn model from training data (model training)
  * Evaluate model on evaluation data (model evaluation)
  * Repeat, revising features
* with production data
  * Evaluate model on production data; monitor (model monitoring)
  * Select production data for retaining (model training + evaluation)
  * Update model regularly (model deployment)

### Feature Engineering

* Identify parameters of interest that a model may learn on
* Conver data into a useful form
* Normalize data
* Include context
* Remove misleading things

### Data Cleaning

* Removing outliers
* Normalizing data
* Missing values

### Learning

* Build a predictor that best describes an outcome for the observed features

### Evaluation

* Prediction accuracy on learned data vs. unseen data
* For binary predictors: false postives vs. false negatives
* For nemeric predictors: average distance between real and predicted value
* For ranking predictors: top K etc

## ML Componeny Tradeoffs

### Qualities of ML Components

* Accuracy
* Capabilities
* Amount og training data needed
* Inference latency
* Learning latency; incremental learning?
* Model sizze
* Explainable? Robust?

## System Architecture Considerations

### Considerations

* How much data is needed as input for the model?
* How much output data is produced by the model?
* How fast/energy consuming is model execution?
* What latency is needed for the application?
* How big is the model? How often does it need to be updated?
* Cost of operating the model? (distribution + execution)
* Opportunities for telemetry?
* What happens if users are offline?

### Typical Designs

* Static intelligence in the product
  * difficult to update
  * good execution latency
  * cheap/offline operation
  * no telemetry to evaluate and improve
* Client-side intelligence
  * updates costly/slow, out of sync problems
  * complexity in clients
  * offline operation, low execution latency
* Server-centric intelligence
  * latency in model execution (remote calls)
  * easy to update and experiment
  * operation cost
  * no offline operation
* Back-end cached intelligence
  * precomputed common results
  * fast execution, partial offline
  * saves bandwidth, complicated updates
* Hybrid models

### Common Design Strategies

* Message-driven, lazy computation, functional programming
  * asynchronous, message passing style
* Replication, containment, supervision
  * replicate and coordinate isolated components
* Data streams, infinite data, immutable facts
  * streaming technologies, data lakes

## Updating Models

* Models are rarely static outside the lab
* Data drift, feedback loops, new features, new requirements

## Planning for Mistakes

### Mistakes Will Happen

* No specification
* ML components detect patterns from data (real and spurious)
* Predictions are often accurate, but mistakes always possible
* Mistakes are not predicable or explainable or similar to human mistakes
* Plan for mistakes

### How Models Can Break

* System outage
* Model outage
* Model errors
* Model degradation

### Mitigating Mistakes

* Investigating in ML
  * e.g., more training data, better data, better features
* Less forceful experience
  * e.g., prompt rather than automate decisions
* Adjust learning parameters
  * e.g., more frequent updates, manual adjustments
* Guardrails
  * e.g., heuristics and constraints on outputs
* Override errors
  * e.g., hardcode specific results

### Telemetry

* Monitor operation/success(accuracy)
* Improve models over time (e.g., detect new features)
* Too much data - sample, summarization, adjustable
* Hard to measure - intended outcome not observable?