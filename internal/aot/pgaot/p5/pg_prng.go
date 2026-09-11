package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_prng_fseed(m *base.Module, l0 int32, l1 float64) {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v31 int64
	_ = v31
	var v36 int64
	_ = v36
	var v41 int64
	_ = v41
	v7 = base.F64_mul(l1, float64(4.503599627370495e+15))
	if base.F64_lt(base.F64_abs(v7), float64(9.223372036854776e+18)) != 0 {
		v11 = base.I64_trunc_f64_s(v7)
		v13 = v11
	} else {
		v13 = int64(-9223372036854775807 - 1)
	}
	v15 = v13 + int64(4354685564936845354)
	v16 = int64(30)
	v19 = int64(-4658895280553007687)
	v20 = (int64(base.Ui64(v15)>>(uint(v16)%64)) ^ v15) * v19
	v21 = int64(27)
	v24 = int64(-7723592293110705685)
	v25 = (int64(base.Ui64(v20)>>(uint(v21)%64)) ^ v20) * v24
	v26 = int64(31)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(base.Ui64(v25)>>(uint(v26)%64)) ^ v25
	v31 = v13 - int64(7046029254386353131)
	v36 = (int64(base.Ui64(v31)>>(uint(v16)%64)) ^ v31) * v19
	v41 = (int64(base.Ui64(v36)>>(uint(v21)%64)) ^ v36) * v24
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(base.Ui64(v41)>>(uint(v26)%64)) ^ v41
	if v31|v15 == int64(0) {
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(1442695040888963407)
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(6364136223846793005)
	} else {
	}
	return
}
func F_pg_prng_seed(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v22 int64
	_ = v22
	var v27 int64
	_ = v27
	var v32 int64
	_ = v32
	v6 = l1 + int64(4354685564936845354)
	v7 = int64(30)
	v10 = int64(-4658895280553007687)
	v11 = (int64(base.Ui64(v6)>>(uint(v7)%64)) ^ v6) * v10
	v12 = int64(27)
	v15 = int64(-7723592293110705685)
	v16 = (int64(base.Ui64(v11)>>(uint(v12)%64)) ^ v11) * v15
	v17 = int64(31)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(base.Ui64(v16)>>(uint(v17)%64)) ^ v16
	v22 = l1 - int64(7046029254386353131)
	v27 = (int64(base.Ui64(v22)>>(uint(v7)%64)) ^ v22) * v10
	v32 = (int64(base.Ui64(v27)>>(uint(v12)%64)) ^ v27) * v15
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(base.Ui64(v32)>>(uint(v17)%64)) ^ v32
	if v22|v6 == int64(0) {
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(1442695040888963407)
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(6364136223846793005)
	} else {
	}
	return
}
