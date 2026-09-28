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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
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
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int64
	_ = v75
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2 < v18 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v32 = v22 * v18 << (uint(int32(1)) % 32)
	v33 = v2
	goto L4
L2:
	;
	v89 = v18
	goto L3
L3:
	;
	v99 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v89
	v111 = F_construct_md_array(m, v17, v16, v99, v14+int32(12), v14+int32(8), int32(25), int32(-1), int32(0), int32(105))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L12
	} else {
		goto L17
	}
L4:
	;
	v37 = int32(1)
	v38 = int64(0)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v42 = v39 + v32<<(uint(int32(2))%32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v43 < int32(0) {
		v73 = v37
		v75 = v38
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v89 = v86
	goto L3
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17+v33<<(uint(int32(3))%32)))) = v75
	*(*uint8)(unsafe.Add(mBase, uint32(v33+v16))) = uint8(v73)
	v85 = v33 + int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v85 < v86 {
		v32 = v32 + int32(2)
		v33 = v85
		goto L4
	} else {
		goto L16
	}
L7:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v46 < int32(0) {
		v73 = v37
		v75 = v38
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v21 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v54 = F_pg_wchar2mb_with_len(m, v49+v43<<(uint(int32(2))%32), v21, v46-v43)
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
	v62 = int32(0)
	v65 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
	v71 = F_DirectFunctionCall3Coll(m, int32(1689), v62, v65, base.I64_extend_i32_s(v43+int32(1)), base.I64_extend_i32_s(v46-v43))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L15
	}
L12:
	;
	return int32(0)
L13:
	;
	v58 = F_cstring_to_text_with_len(m, v21, v54)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v73 = int32(0)
	v75 = base.I64_extend_i32_u(v58)
	goto L6
L15:
	;
	v73 = v62
	v75 = v71
	goto L6
L16:
	;
	goto L5
L17:
	;
	m.G0 = v14 + int32(16)
	return v111
}
func F_regexp_substr_no_start(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_regexp_substr(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
