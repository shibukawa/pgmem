package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cnt_sml(m *base.Module, l0 int32, l1 int32, l2 int32) float32 {
	mBase := m.M
	_ = mBase
	var v10 float32
	_ = v10
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v91 float32
	_ = v91
	v10 = float32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = int32(2)
	v14 = int32(5)
	v15 = int32(base.Ui32(v11)>>(uint(v12)%32)) - v14
	v16 = int32(3)
	v17 = base.I32_div_u_s(v15, v16)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = int32(base.Ui32(v18)>>(uint(v12)%32)) - v14
	v24 = base.I32_div_u_s(v22, v16)
	if base.Ui32(v22) < base.Ui32(v16) {
		v91 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v91
L2:
	;
	if base.Ui32(v15) < base.Ui32(int32(3)) {
		v91 = v10
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = int32(5)
	v30 = l0 + v29
	v32 = l1 + v29
	v33 = v30
	v34 = v32
	v36 = int32(0)
	goto L4
L4:
	;
	v45 = base.I32_div_s(v34-v32, int32(3))
	if v45 < v17 {
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
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_cnt_sml[0]))
	v49 = m.T0[v48].(func(*base.Module, int32, int32) int32)(m, v33, v34)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v74 = v36
	goto L8
L8:
	;
	goto L5
L9:
	;
	v70 = base.I32_div_s(v65-v30, int32(3))
	if v70 < v24 {
		v33 = v65
		v34 = v66
		v36 = v67
		goto L4
	} else {
		goto L18
	}
L10:
	;
	return float32(0)
L11:
	;
	if int32(0) <= v49 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v49 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v61 = v34
	v62 = v36
	goto L14
L14:
	;
	v65 = v33 + int32(3)
	v66 = v61
	v67 = v62
	goto L9
L15:
	;
	v65 = v33
	v66 = v34 + int32(3)
	v67 = v36
	goto L9
L16:
	;
	goto L17
L17:
	;
	v61 = v34 + int32(3)
	v62 = v36 + int32(1)
	goto L14
L18:
	;
	v74 = v67
	goto L8
L19:
	;
	v78 = v74
	goto L21
L20:
	;
	v78 = v17
	goto L21
L21:
	;
	v91 = base.F32_div(base.F32_convert_i32_s(v74), base.F32_convert_i32_s(v24-v74+v78))
	goto L1
}
