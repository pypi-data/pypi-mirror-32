# kumlog
Kumparan Logging Json Format

# Example Usage

```
def main():
    logger.debug("test DEBUG logger")
    logger.info("test INFO logger")
    logger.warning("test WARNING logger")
    logger.error("test ERROR logger")
    logger.critical("test CRITICAL logger")

    try:
        raise ValueError("Input is not valid")
    except Exception as e:
        logger.exception(e)

    try:
        raise OSError("not found")
    except Exception as e:
        logger.exception(e)


if __name__ == '__main__':
    # Setup the logging handler and the formatter
    handler = logging.StreamHandler(stream=sys.stdout)
    formatter = StackdriverLoggingFormatter(service_name="stackdriver-logging",
                                            service_version="latest", environment="staging")
    handler.setFormatter(formatter)
    logger.addHandler(handler)
    logger.setLevel(logging.DEBUG)

    main()

```
