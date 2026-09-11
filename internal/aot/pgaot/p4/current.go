package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetCurrentTransactionId(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v5 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v4))))
	if v5 == int64(0) {
		F_AssignTransactionId(m, v4)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
			v13 = v12
			return base.I32_wrap_i64(v13)
		}
	} else {
		v13 = v5
		return base.I32_wrap_i64(v13)
	}
}
func F_current_schemas(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_fetch_search_path(m, base.B2i32(v6 != v2))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_list_free(m, v9)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L18
	}
L2:
	;
	return int32(0)
L3:
	;
	if v9 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = F_palloc(m, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v21 = F_palloc(m, v18<<(uint(int32(2))%32))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v54 = v2
	v56 = v16
	goto L1
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v23 <= int32(0) {
		v54 = v2
		v56 = v21
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v27 = v2
	v28 = v2
	goto L10
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v27<<(uint(int32(2))%32))))
	v36 = F_get_namespace_name(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L12
	}
L11:
	;
	v54 = v48
	v56 = v21
	goto L1
L12:
	;
	if v36 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v43 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v36)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L16
	}
L14:
	;
	v48 = v28
	goto L15
L15:
	;
	v50 = v27 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v50 < v51 {
		v27 = v50
		v28 = v48
		goto L10
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+v28<<(uint(int32(2))%32)))) = v43
	v48 = v28 + int32(1)
	goto L15
L17:
	;
	goto L11
L18:
	;
	v61 = F_construct_array_builtin(m, v56, v54, int32(19))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	return v61
}
