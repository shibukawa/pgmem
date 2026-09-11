package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SN_create_env(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	v7 = F_palloc0(m, int32(32))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	if v70 != 0 {
		goto L24
	} else {
		goto L25
	}
L2:
	;
	return int32(0)
L3:
	;
	if v7 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v11 = F_create_s(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	v63 = int32(0)
	goto L6
L6:
	;
	return v63
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v11
	if v11 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if l0 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if l1 != 0 {
		goto L19
	} else {
		goto L20
	}
L10:
	;
	v20 = F_palloc0(m, l0<<(uint(int32(2))%32))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v20
	if v20 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v25 = int32(0)
	if l0 <= v25 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v30 = v25
	goto L14
L14:
	;
	v33 = F_create_s(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L16
	}
L15:
	;
	goto L9
L16:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v35+v30<<(uint(int32(2))%32)))) = v33
	if v33 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v43 = v30 + int32(1)
	if v43 != l0 {
		v30 = v43
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v52 = F_palloc0(m, l1<<(uint(int32(2))%32))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v63 = v7
	goto L6
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v52
	if v52 == int32(0) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	if int32(0) < l0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	F_pfree(m, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L35
	}
L27:
	;
	v76 = int32(0)
	goto L30
L28:
	;
	v95 = v70
	goto L29
L29:
	;
	F_pfree(m, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L2
	} else {
		goto L34
	}
L30:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v76<<(uint(int32(2))%32))))
	F_lose_s(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L32
	}
L31:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	v95 = v89
	goto L29
L32:
	;
	v87 = v76 + int32(1)
	if v87 != l0 {
		v76 = v87
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L26
L35:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v106 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	F_lose_s(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	F_pfree(m, v7)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	return int32(0)
}
func F_sn_array_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v8 <= v5+int32(1) {
		F_appendStringInfoChar(m, v4, int32(91))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v19 = int32(91)
		*(*uint8)(unsafe.Add(mBase, uint32(v17+v5))) = uint8(v19)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
		v24 = v22 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v24
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		v28 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v26+v24))) = uint8(v28)
		return v28
	}
}
func F_sn_object_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v8 <= v5+int32(1) {
		F_appendStringInfoChar(m, v4, int32(123))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v19 = int32(123)
		*(*uint8)(unsafe.Add(mBase, uint32(v17+v5))) = uint8(v19)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
		v24 = v22 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v24
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		v28 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v26+v24))) = uint8(v28)
		return v28
	}
}
