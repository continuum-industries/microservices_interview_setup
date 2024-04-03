import pathlib
import sys

if (path := str(pathlib.Path(__file__).parent.resolve())) not in sys.path:
    sys.path.append(path)


CALCULATOR_HANDLER = "calculator.proto"
