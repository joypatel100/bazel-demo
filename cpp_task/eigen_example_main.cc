#include <iostream>
#include <Eigen/Dense>

// If downloaded Eigen raw:
// g++ -I /Users/joy/Downloads/eigen-eigen-5a0156e40feb/ eigen_code.cc -o eigen_code
//
// If installed with opencv via homebew:
// g++ `pkg-config --cflags --libs eigen3` eigen_code.cc -o eigen_code


using Eigen::MatrixXd;
int main()
{
  MatrixXd m(2,2);
  m(0,0) = 3;
  m(1,0) = 2.5;
  m(0,1) = -1;
  m(1,1) = m(1,0) + m(0,1);
  std::cout << m << std::endl;
}
