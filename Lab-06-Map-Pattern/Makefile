# Joseph Kehoe 13/9/2017
# Modified for macOS OpenMP

CXX = clang++

OMP_FLAGS = -Xpreprocessor -fopenmp
OMP_INC = -I$(shell brew --prefix libomp)/include
OMP_LIB = -L$(shell brew --prefix libomp)/lib -lomp

CXXFLAGS = -std=c++11 -pthread $(OMP_FLAGS) $(OMP_INC)
LDFLAGS = $(OMP_LIB)

SAXPY:
	$(CXX) $(CXXFLAGS) saxpy.cpp $(LDFLAGS) -o saxpy.out

MANDEL:
	$(CXX) $(CXXFLAGS) mandelbroit.cpp $(LDFLAGS) -o mandelbroit.out

clean:
	rm -f saxpy.out mandelbroit.out