SDK_REPO=$1
BUILD_DIR=$2
TESTS_DIR=$3

echo "SDK_REPO=$SDK_REPO"
echo "BUILD_DIR=$BUILD_DIR"
echo "TESTS_DIR=$TESTS_DIR"

# Copy license
cp $SDK_REPO/LICENSE $BUILD_DIR/License.txt

# Copy python config file
cp $TESTS_DIR/setup.cfg $BUILD_DIR/setup.cfg

# Compile module
cd $BUILD_DIR
python setup.py build

# Run tests
echo "Running tests"
cd "$TESTS_DIR"
echo "Install requests..."
python -m pip install requests
echo "Install six..."
python -m pip install six
echo "Run unit tests"
python -m unittest SdkTests
