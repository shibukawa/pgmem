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
	var v23 float64
	_ = v23
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v29 float64
	_ = v29
	var v31 float64
	_ = v31
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
	var v52 int32
	_ = v52
	var v53 float64
	_ = v53
	var v55 int32
	_ = v55
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v65 float64
	_ = v65
	var v73 int32
	_ = v73
	var v74 float64
	_ = v74
	var v76 float64
	_ = v76
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
		v23 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v9-int32(-64)))) = v23
		v25 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		*(*float64)(unsafe.Add(mBase, uint32(v9)+72)) = v25
		v27 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v9)+80)) = v27
		v29 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		*(*float64)(unsafe.Add(mBase, uint32(v9)+88)) = v29
		v31 = *(*float64)(unsafe.Add(mBase, uint32(v7)+24))
		*(*float64)(unsafe.Add(mBase, uint32(v9)+96)) = v31
		v34 = v7 + int32(24)
		v35 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		if base.Ui64(base.I64_reinterpret_f64(v35)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			v42 = v7 + int32(16)
			v43 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			if base.F64_gt(v43, v35) != 0 {
				v52 = v42
				v53 = v43
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v43)&int64(9223372036854775807)) {
					v52 = v42
					v53 = v43
				} else {
					v52 = v7
					v53 = v35
				}
			}
		} else {
			v52 = v7
			v53 = v35
		}
		v55 = v7 + int32(8)
		*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v53
		v57 = *(*float64)(unsafe.Add(mBase, uint32(v52)))
		*(*float64)(unsafe.Add(mBase, uint32(v9)+24)) = v57
		v59 = *(*float64)(unsafe.Add(mBase, uint32(v34)))
		if base.Ui64(base.I64_reinterpret_f64(v59)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			v65 = *(*float64)(unsafe.Add(mBase, uint32(v55)))
			if base.F64_gt(v65, v59) != 0 {
				v73 = v34
				v74 = v65
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v65)&int64(9223372036854775807)) {
					v73 = v34
					v74 = v65
				} else {
					v73 = v55
					v74 = v59
				}
			}
		} else {
			v73 = v55
			v74 = v59
		}
		*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v74
		v76 = *(*float64)(unsafe.Add(mBase, uint32(v73)))
		*(*float64)(unsafe.Add(mBase, uint32(v9)+32)) = v76
		return v9
	}
}
