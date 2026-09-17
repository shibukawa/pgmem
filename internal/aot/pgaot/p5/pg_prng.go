package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_prng_fseed(m *base.Module, l0 int32, l1 float64) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v25 int64
	_ = v25
	var v30 int64
	_ = v30
	var v35 int64
	_ = v35
	v7 = base.I64_trunc_sat_f64_s(base.F64_mul(l1, float64(4.503599627370495e+15)))
	v9 = v7 + int64(4354685564936845354)
	v10 = int64(30)
	v13 = int64(-4658895280553007687)
	v14 = (int64(base.Ui64(v9)>>(uint(v10)%64)) ^ v9) * v13
	v15 = int64(27)
	v18 = int64(-7723592293110705685)
	v19 = (int64(base.Ui64(v14)>>(uint(v15)%64)) ^ v14) * v18
	v20 = int64(31)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(base.Ui64(v19)>>(uint(v20)%64)) ^ v19
	v25 = v7 - int64(7046029254386353131)
	v30 = (int64(base.Ui64(v25)>>(uint(v10)%64)) ^ v25) * v13
	v35 = (int64(base.Ui64(v30)>>(uint(v15)%64)) ^ v30) * v18
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(base.Ui64(v35)>>(uint(v20)%64)) ^ v35
	return
}
func F_pg_prng_seed(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v21 int64
	_ = v21
	var v26 int64
	_ = v26
	var v31 int64
	_ = v31
	v5 = l1 + int64(4354685564936845354)
	v6 = int64(30)
	v9 = int64(-4658895280553007687)
	v10 = (int64(base.Ui64(v5)>>(uint(v6)%64)) ^ v5) * v9
	v11 = int64(27)
	v14 = int64(-7723592293110705685)
	v15 = (int64(base.Ui64(v10)>>(uint(v11)%64)) ^ v10) * v14
	v16 = int64(31)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(base.Ui64(v15)>>(uint(v16)%64)) ^ v15
	v21 = l1 - int64(7046029254386353131)
	v26 = (int64(base.Ui64(v21)>>(uint(v6)%64)) ^ v21) * v9
	v31 = (int64(base.Ui64(v26)>>(uint(v11)%64)) ^ v26) * v14
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(base.Ui64(v31)>>(uint(v16)%64)) ^ v31
	return
}
