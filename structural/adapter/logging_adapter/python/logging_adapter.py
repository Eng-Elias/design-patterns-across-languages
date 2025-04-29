import abc

# --- Adaptee --- (The incompatible third-party library)
class ThirdPartyLogger:
    """A hypothetical third-party logger with an incompatible interface."""
    def record(self, severity: str, message: str):
        print(f"[3rdPartyLogger - {severity.upper()}]: {message}")

# --- Target --- (The interface the client code expects)
class Logger(abc.ABC):
    """The interface our application expects for logging."""
    @abc.abstractmethod
    def log_info(self, message: str):
        pass

    @abc.abstractmethod
    def log_warning(self, message: str):
        pass

    @abc.abstractmethod
    def log_error(self, message: str):
        pass

# --- Adapter ---
class LoggerAdapter(Logger):
    """Adapts the ThirdPartyLogger to the Logger interface."""
    def __init__(self, adaptee: ThirdPartyLogger):
        self._adaptee = adaptee

    def log_info(self, message: str):
        self._adaptee.record("info", message)

    def log_warning(self, message: str):
        self._adaptee.record("warning", message)

    def log_error(self, message: str):
        self._adaptee.record("error", message)

# --- Client Code ---
class ApplicationService:
    """A service that uses the Logger interface."""
    def __init__(self, logger: Logger):
        self._logger = logger

    def perform_operation(self, data: str):
        self._logger.log_info(f"Starting operation with data: {data}")
        try:
            # Simulate an operation
            if not data:
                raise ValueError("Data cannot be empty")
            if len(data) < 5:
                 self._logger.log_warning(f"Data '{data}' is quite short.")
            # ... perform actual operation ...
            self._logger.log_info("Operation completed successfully.")
        except Exception as e:
            self._logger.log_error(f"Operation failed: {e}")
