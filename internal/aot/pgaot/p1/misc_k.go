package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_KnownAssignedXidsRemove(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v17 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
	F_errmsg_internal(m, int32(54801), v13)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[666]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v33 = v31 - int32(1)
	if v33 < v30 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_errfinish(m, int32(490691), int32(4991), int32(341491))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	m.G0 = v13 + int32(16)
	return
L9:
	;
	v38 = v33
	v39 = v30
	goto L10
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[777]))
	v47 = v38 + v39
	v48 = int32(2)
	v49 = base.I32_div_s(v47, v48)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46+v49<<(uint(v48)%32))))
	if v53 != l0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v47 < int32(-1) {
		goto L8
	} else {
		goto L26
	}
L12:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v53))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	goto L11
L15:
	;
	if v68 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v68 = base.B2i32(base.Ui32(l0) < base.Ui32(v53))
	goto L15
L17:
	;
	goto L18
L18:
	;
	v68 = int32(base.Ui32(l0-v53) >> (uint(int32(31)) % 32))
	goto L15
L19:
	;
	v69 = v39
	goto L21
L20:
	;
	v69 = v49 + int32(1)
	goto L21
L21:
	;
	if v68 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v72 = v49 - int32(1)
	goto L24
L23:
	;
	v72 = v38
	goto L24
L24:
	;
	if v69 <= v72 {
		v38 = v72
		v39 = v69
		goto L10
	} else {
		goto L25
	}
L25:
	;
	goto L8
L26:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[778]))
	v78 = v77 + v49
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v79 != int32(1) {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v82 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v82)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v84 - int32(1)
	if v30 != v49 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v90 = v30
	goto L30
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v109
	goto L8
L30:
	;
	v100 = v90 + int32(1)
	if v100 < v31 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v106 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v106
	v109 = v106
	goto L29
L32:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+v77))))
	if v103 == int32(0) {
		v90 = v100
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	goto L31
L35:
	;
	v109 = v100
	goto L29
}
func F_koi8r_to_mic(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(22), int32(7))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v21 = F_latin2mic(m, v6, v5, v10, int32(139), int32(22), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			return v21
		}
	}
}
