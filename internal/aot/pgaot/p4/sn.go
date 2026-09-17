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
	var v31 int32
	_ = v31
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
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
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
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	if v71 != 0 {
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
	v64 = int32(0)
	goto L6
L6:
	;
	return v64
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
	v31 = v25
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
	*(*int32)(unsafe.Add(mBase, uint32(v35+v31<<(uint(int32(2))%32)))) = v33
	if v33 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v43 = v31 + int32(1)
	if v43 != l0 {
		v31 = v43
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
	v64 = v7
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
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	F_pfree(m, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L35
	}
L27:
	;
	v78 = int32(0)
	goto L30
L28:
	;
	v96 = v71
	goto L29
L29:
	;
	F_pfree(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L2
	} else {
		goto L34
	}
L30:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+v78<<(uint(int32(2))%32))))
	F_lose_s(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L32
	}
L31:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	v96 = v90
	goto L29
L32:
	;
	v88 = v78 + int32(1)
	if v88 != l0 {
		v78 = v88
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
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v107 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	F_lose_s(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
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
	v111 = m.ExcPending
	if v111 != 0 {
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
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13993(m, l0, int32(91))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_sn_object_start(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13993(m, l0, int32(123))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
