from logging_adapter import ThirdPartyLogger, LoggerAdapter, ApplicationService

def main():
    """Demonstrates using the Adapter pattern."""
    print("--- Using the Adapter for the Third-Party Logger ---")

    # Create the Adaptee (the incompatible logger)
    third_party_logger = ThirdPartyLogger()

    # Create the Adapter, wrapping the Adaptee
    logger_adapter = LoggerAdapter(third_party_logger)

    # The client code (ApplicationService) uses the standard Logger interface
    # It doesn't know it's talking to an adapter or a third-party logger.
    app_service = ApplicationService(logger_adapter)

    print("\nPerforming operations:")
    app_service.perform_operation("ImportantData123")
    print("---")
    app_service.perform_operation("abc") # Should trigger a warning
    print("---")
    app_service.perform_operation("")    # Should trigger an error

if __name__ == "__main__":
    main()
