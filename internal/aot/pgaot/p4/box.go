package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_box_poly(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 float64
	_ = v15
	var v17 float64
	_ = v17
	var v19 float64
	_ = v19
	var v21 float64
	_ = v21
	var v23 float64
	_ = v23
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v29 float64
	_ = v29
	var v32 int32
	_ = v32
	var v33 float64
	_ = v33
	var v41 float64
	_ = v41
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v63 float64
	_ = v63
	var v72 int32
	_ = v72
	var v73 float64
	_ = v73
	var v75 float64
	_ = v75
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_palloc(m, int32(104))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(17179869600)
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		*(*float64)(unsafe.Add(mBase, uint32(v9)+40)) = v15
		v17 = *(*float64)(unsafe.Add(mBase, uint32(v7)+24))
		*(*float64)(unsafe.Add(mBase, uint32(v9)+48)) = v17
		v19 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		*(*float64)(unsafe.Add(mBase, uint32(v9)+56)) = v19
		v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v9)+64)) = v21
		v23 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		*(*float64)(unsafe.Add(mBase, uint32(v9)+72)) = v23
		v25 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v9)+80)) = v25
		v27 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		*(*float64)(unsafe.Add(mBase, uint32(v9)+88)) = v27
		v29 = *(*float64)(unsafe.Add(mBase, uint32(v7)+24))
		*(*float64)(unsafe.Add(mBase, uint32(v9)+96)) = v29
		v32 = v7 + int32(24)
		v33 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		if base.Ui64(base.I64_reinterpret_f64(v33)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			v41 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			if base.F64_gt(v41, v33)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v41)&int64(9223372036854775807))) != 0 {
				v50 = v41
				v51 = v7 + int32(16)
			} else {
				v50 = v33
				v51 = v7
			}
		} else {
			v50 = v33
			v51 = v7
		}
		v53 = v7 + int32(8)
		*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v50
		v55 = *(*float64)(unsafe.Add(mBase, uint32(v51)))
		*(*float64)(unsafe.Add(mBase, uint32(v9)+24)) = v55
		v57 = *(*float64)(unsafe.Add(mBase, uint32(v32)))
		if base.Ui64(base.I64_reinterpret_f64(v57)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			v63 = *(*float64)(unsafe.Add(mBase, uint32(v53)))
			if base.F64_gt(v63, v57)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v63)&int64(9223372036854775807))) != 0 {
				v72 = v32
				v73 = v63
			} else {
				v72 = v53
				v73 = v57
			}
		} else {
			v72 = v53
			v73 = v57
		}
		*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v73
		v75 = *(*float64)(unsafe.Add(mBase, uint32(v72)))
		*(*float64)(unsafe.Add(mBase, uint32(v9)+32)) = v75
		return v9
	}
}
