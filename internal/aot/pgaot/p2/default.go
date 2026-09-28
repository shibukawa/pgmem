package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FindDefaultConversionProc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v117 int32
	_ = v117
	v3 = int32(0)
	F_recomputeNamespacePath(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_FindDefaultConversionProc[0]))
	if v15 == int32(0) {
		v117 = v3
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v117
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if int32(0) < v18 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v27 = v3
	goto L8
L6:
	;
	goto L7
L7:
	;
	v117 = int32(0)
	goto L3
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v27<<(uint(int32(2))%32))))
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_FindDefaultConversionProc[1]))
	if v34 != v36 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v43 = F_SearchSysCacheList(m, int32(17), int32(3), base.I64_extend_i32_u(v34), base.I64_extend_i32_s(l0), base.I64_extend_i32_s(l1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v102 = v27 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v102 < v103 {
		v27 = v102
		goto L8
	} else {
		goto L27
	}
L13:
	;
	if v91 != 0 {
		v117 = v91
		goto L3
	} else {
		goto L26
	}
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	if v45 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_ReleaseCatCacheList(m, v43)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v56 = int32(0)
	goto L20
L18:
	;
	v91 = int32(0)
	goto L13
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v69)+84))
	F_ReleaseCatCacheList(m, v43)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L25
	}
L20:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v43-int32(-64)+v56<<(uint(int32(2))%32))))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
	v69 = v67 + v68
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+88)))
	if v70 == int32(1) {
		goto L19
	} else {
		goto L22
	}
L21:
	;
	F_ReleaseCatCacheList(m, v43)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	v74 = v56 + int32(1)
	if v74 != v45 {
		v56 = v74
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v91 = int32(0)
	goto L13
L25:
	;
	v91 = v79
	goto L13
L26:
	;
	goto L12
L27:
	;
	goto L9
}
