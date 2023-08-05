#ifndef THC_NUMERICS_INC
#define THC_NUMERICS_INC

#include <THC/THCHalf.h>

#ifdef CUDA_HALF_TENSOR
#ifdef __CUDA_ARCH__
#define h2f(A) __half2float(A)
#define f2h(A) __float2half(A)
#else  // CUDA_ARCH__
#define h2f(A) THC_half2float(A)
#define f2h(A) THC_float2half(A)
#endif  // CUDA_ARCH__
#endif  // CUDA_HALF_TENSOR

template<typename T>
struct THCNumerics {
  static inline __host__ __device__ T add(T a, T b) { return a + b; }
  static inline __host__ __device__ T sub(T a, T b) { return a - b; }
  static inline __host__ __device__ T mul(T a, T b) { return a * b; }
  static inline __host__ __device__ T div(T a, T b) { return a / b; }
  static inline __host__ __device__ T neg(T a) { return -a; }
};

#ifdef CUDA_HALF_TENSOR
template<>
struct THCNumerics<half> {
  static inline __host__ __device__ half add(half a, half b) { return f2h(h2f(a) + h2f(b)); }
  static inline __host__ __device__ half sub(half a, half b) { return f2h(h2f(a) - h2f(b)); }
  static inline __host__ __device__ half mul(half a, half b) { return f2h(h2f(a) * h2f(b)); }
  static inline __host__ __device__ half div(half a, half b) { return f2h(h2f(a) / h2f(b)); }
  static inline __host__ __device__ half neg(half a) { return f2h(-h2f(a)); }
};
#endif // CUDA_HALF_TENSOR

template <typename In, typename Out>
struct ScalarConvert {
  static __host__ __device__ Out to(const In v) { return (Out) v; }
};

#ifdef CUDA_HALF_TENSOR
template <typename Out>
struct ScalarConvert<half, Out> {
  static __host__ __device__ Out to(const half v) { return (Out) h2f(v); }
};

template <typename In>
struct ScalarConvert<In, half> {
  static __host__ __device__ half to(const In v) { return f2h((float) v); }
};
#endif  // CUDA_HALF_TENSOR

#endif // THC_NUMERICS_INC
