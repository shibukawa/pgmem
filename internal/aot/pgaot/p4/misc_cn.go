package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cnt_sml(m *base.Module, l0 int32, l1 int32, l2 int32) float32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v94 float32
	_ = v94
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = int32(2)
	v14 = int32(5)
	v15 = int32(base.Ui32(v11)>>(uint(v12)%32)) - v14
	v16 = int32(3)
	v17 = base.I32_div_u_s(v15, v16)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = int32(base.Ui32(v18)>>(uint(v12)%32)) - v14
	v24 = base.I32_div_u_s(v22, v16)
	if base.B2i32(base.Ui32(v22) < base.Ui32(v16))|base.B2i32(base.Ui32(v15) < base.Ui32(v16)) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v94 = float32(0)
	goto L3
L2:
	;
	v31 = int32(5)
	v32 = l0 + v31
	v34 = l1 + v31
	v35 = v32
	v36 = v34
	v38 = int32(0)
	goto L4
L3:
	;
	return v94
L4:
	;
	v47 = base.I32_div_s(v36-v34, int32(3))
	if v47 < v17 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if l2 != 0 {
		goto L19
	} else {
		goto L20
	}
L6:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_cnt_sml[0]))
	v51 = m.T0[v50].(func(*base.Module, int32, int32) int32)(m, v35, v36)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v76 = v38
	goto L8
L8:
	;
	goto L5
L9:
	;
	v72 = base.I32_div_s(v67-v32, int32(3))
	if v72 < v24 {
		v35 = v67
		v36 = v68
		v38 = v69
		goto L4
	} else {
		goto L18
	}
L10:
	;
	return float32(0)
L11:
	;
	if int32(0) <= v51 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v51 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v63 = v36
	v64 = v38
	goto L14
L14:
	;
	v67 = v35 + int32(3)
	v68 = v63
	v69 = v64
	goto L9
L15:
	;
	v67 = v35
	v68 = v36 + int32(3)
	v69 = v38
	goto L9
L16:
	;
	goto L17
L17:
	;
	v63 = v36 + int32(3)
	v64 = v38 + int32(1)
	goto L14
L18:
	;
	v76 = v69
	goto L8
L19:
	;
	v80 = v76
	goto L21
L20:
	;
	v80 = v17
	goto L21
L21:
	;
	v94 = base.F32_div(base.F32_convert_i32_s(v76), base.F32_convert_i32_s(v24-v76+v80))
	goto L3
}
