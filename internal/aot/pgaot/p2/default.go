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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v114 int32
	_ = v114
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
		v114 = v3
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v114
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
	v114 = int32(0)
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
	v40 = F_SearchSysCacheList(m, int32(17), int32(3), v34, l0, l1)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v99 = v27 + int32(1)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v99 < v100 {
		v27 = v99
		goto L8
	} else {
		goto L27
	}
L13:
	;
	if v88 != 0 {
		v114 = v88
		goto L3
	} else {
		goto L26
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if v42 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_ReleaseCatCacheList(m, v40)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v53 = int32(0)
	goto L20
L18:
	;
	v88 = int32(0)
	goto L13
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v66)+84))
	F_ReleaseCatCacheList(m, v40)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L25
	}
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(48)+v53<<(uint(int32(2))%32))))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+56))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+22)))
	v66 = v64 + v65
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+88)))
	if v67 == int32(1) {
		goto L19
	} else {
		goto L22
	}
L21:
	;
	F_ReleaseCatCacheList(m, v40)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	v71 = v53 + int32(1)
	if v71 != v42 {
		v53 = v71
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v88 = int32(0)
	goto L13
L25:
	;
	v88 = v76
	goto L13
L26:
	;
	goto L12
L27:
	;
	goto L9
}
