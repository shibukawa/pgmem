package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dist_lp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = F_line_closept_point(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_Float8GetDatum(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_dist_ppoly(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_dist_ppoly_internal(m, v2, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = F_Float8GetDatum(m, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	}
}
func F_dist_ppoly_internal(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 float64
	_ = v20
	var v22 float64
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v34 float64
	_ = v34
	var v36 int32
	_ = v36
	var v38 float64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v51 float64
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 float64
	_ = v56
	var v58 float64
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 float64
	_ = v65
	var v67 float64
	_ = v67
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v78 float64
	_ = v78
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v85 int32
	_ = v85
	var v95 float64
	_ = v95
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v15 = l1 + int32(40)
	v16 = F_point_inside(m, l0, v13, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return v95
L2:
	;
	return float64(0)
L3:
	;
	if v16 != 0 {
		v95 = float64(0)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = *(*float64)(unsafe.Add(mBase, uint32(l1)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v11))) = v20
	v22 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v27 = v15 + v24<<(uint(int32(4))%32)
	v30 = *(*float64)(unsafe.Add(mBase, uint32(v27-int32(16))))
	*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v30
	v34 = *(*float64)(unsafe.Add(mBase, uint32(v27-int32(8))))
	*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v34
	v36 = int32(0)
	v38 = F_lseg_closept_point(m, v36, v11, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v40-int32(1) <= int32(0) {
		v95 = v38
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v48 = v36
	v51 = v38
	goto L7
L7:
	;
	v53 = int32(4)
	v55 = v15 + v48<<(uint(v53)%32)
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v55)))
	*(*float64)(unsafe.Add(mBase, uint32(v11))) = v56
	v58 = *(*float64)(unsafe.Add(mBase, uint32(v55)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v58
	v61 = v48 + int32(1)
	v64 = v15 + v61<<(uint(v53)%32)
	v65 = *(*float64)(unsafe.Add(mBase, uint32(v64)))
	*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v65
	v67 = *(*float64)(unsafe.Add(mBase, uint32(v64)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v67
	v70 = F_lseg_closept_point(m, int32(0), v11, l0)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L9
	}
L8:
	;
	v95 = v84
	goto L1
L9:
	;
	if base.Ui64(base.I64_reinterpret_f64(v70)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if base.F64_gt(v51, v70) != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v84 = v51
	goto L12
L12:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v61 < v85-int32(1) {
		v48 = v61
		v51 = v84
		goto L7
	} else {
		goto L19
	}
L13:
	;
	v78 = v70
	goto L15
L14:
	;
	v78 = v51
	goto L15
L15:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(base.F64_abs(v51))) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v83 = v70
	goto L18
L17:
	;
	v83 = v78
	goto L18
L18:
	;
	v84 = v83
	goto L12
L19:
	;
	goto L8
}
