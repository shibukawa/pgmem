package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_build_regexp_match_result(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2 < v17 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v27 = v2
	v32 = v21 * v17 << (uint(int32(1)) % 32)
	goto L4
L2:
	;
	v85 = v17
	goto L3
L3:
	;
	v94 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v85
	v106 = F_construct_md_array(m, v16, v15, v94, v13+int32(12), v13+int32(8), int32(25), int32(-1), int32(0), int32(105))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L12
	} else {
		goto L17
	}
L4:
	;
	v35 = int32(1)
	v36 = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v40 = v37 + v32<<(uint(int32(2))%32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v41 < v36 {
		v69 = v35
		v71 = v36
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v85 = v82
	goto L3
L6:
	;
	v72 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v16+v27<<(uint(v72)%32)))) = v71
	*(*uint8)(unsafe.Add(mBase, uint32(v27+v15))) = uint8(v69)
	v81 = v27 + int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v81 < v82 {
		v27 = v81
		v32 = v32 + v72
		goto L4
	} else {
		goto L16
	}
L7:
	;
	v44 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v45 < v44 {
		v69 = v35
		v71 = v44
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v20 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v54 = F_pg_wchar2mb_with_len(m, v49+v41<<(uint(int32(2))%32), v20, v45-v41)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v60 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v67 = F_DirectFunctionCall3Coll(m, int32(1493), v60, v63, v41+int32(1), v45-v41)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L15
	}
L12:
	;
	return int32(0)
L13:
	;
	v58 = F_cstring_to_text_with_len(m, v20, v54)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v69 = int32(0)
	v71 = v58
	goto L6
L15:
	;
	v69 = v60
	v71 = v67
	goto L6
L16:
	;
	goto L5
L17:
	;
	m.G0 = v13 + int32(16)
	return v106
}
func F_regexp_substr_no_start(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_regexp_substr(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
