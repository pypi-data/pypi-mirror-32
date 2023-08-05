# Import model and loss
from starttf.models.gru_function_classifier import create_model
from starttf.examples.gru_function_classifier.loss import create_loss

# Import utility functions for training and hyper parameter management.
from starttf.estimators.scientific_estimator import easy_train_and_evaluate
from starttf.utils.hyperparams import load_params

if __name__ == "__main__":
    # Load the hyperparameters
    hyper_params = load_params("starttf/examples/gru_function_classifier/hyper_params.json")

    # Invoke the training
    easy_train_and_evaluate(hyper_params, create_model, create_loss)
