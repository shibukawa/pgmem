package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AtEOXact_ComboCid(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[66])) = v2
	*(*int32)(unsafe.Add(mBase, _consts[1246])) = v2
	*(*int32)(unsafe.Add(mBase, _consts[1248])) = v2
	*(*int32)(unsafe.Add(mBase, _consts[1247])) = v2
	return
}
func F_AtEOXact_Inval(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	*(*int32)(unsafe.Add(mBase, _consts[1153])) = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[1154]))
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l0 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1154])) = int32(0)
	goto L3
L5:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+16)))
	if v11 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if v58 < v59 {
		goto L23
	} else {
		goto L24
	}
L8:
	;
	F_RelationCacheInitFilePreInvalidate(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v18 = v10
	goto L10
L10:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v19
	v23 = v18 + int32(32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v28 = v19 - v27
	if int32(0) < v28 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	return
L12:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[1154]))
	v18 = v17
	goto L10
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[1155]))
	F_SendSharedInvalidMessages(m, v32+v27<<(uint(int32(4))%32), v28)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	v39 = v24
	goto L15
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v41 = v39 - v40
	if int32(0) < v41 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v39 = v38
	goto L15
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[1156]))
	F_SendSharedInvalidMessages(m, v45+v40<<(uint(int32(4))%32), v41)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L11
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[1154]))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+16)))
	if v53 != int32(1) {
		goto L4
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	F_RelationCacheInitFilePostInvalidate(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	goto L4
L23:
	;
	v61 = v58
	goto L26
L24:
	;
	goto L25
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	if v82 <= v81 {
		goto L4
	} else {
		goto L30
	}
L26:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[1155]))
	F_LocalExecuteInvalidationMessage(m, v67+v61<<(uint(int32(4))%32))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L11
	} else {
		goto L28
	}
L27:
	;
	goto L25
L28:
	;
	v74 = v61 + int32(1)
	if v74 != v59 {
		v61 = v74
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v84 = v81
	goto L31
L31:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _consts[1156]))
	F_LocalExecuteInvalidationMessage(m, v90+v84<<(uint(int32(4))%32))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L11
	} else {
		goto L33
	}
L32:
	;
	goto L4
L33:
	;
	v97 = v84 + int32(1)
	if v97 != v82 {
		v84 = v97
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
}
func F_AtEOXact_LargeObject(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _consts[466])))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v52 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[464])) = v52
	*(*int32)(unsafe.Add(mBase, _consts[465])) = v52
	v58 = *(*int32)(unsafe.Add(mBase, _consts[467]))
	if v58 != 0 {
		goto L19
	} else {
		goto L20
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _consts[464]))
	if v12 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[465]))
	v18 = v12
	v19 = int32(0)
	v20 = v16
	goto L7
L7:
	;
	v25 = v20 + v19<<(uint(int32(2))%32)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L4
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v29 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v40 = v18
	v41 = v20
	goto L11
L11:
	;
	v43 = v19 + int32(1)
	if v43 < v40 {
		v18 = v40
		v19 = v43
		v20 = v41
		goto L7
	} else {
		goto L18
	}
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[402]))
	F_UnregisterSnapshotFromOwner(m, v29, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	F_pfree(m, v26)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L15
	} else {
		goto L17
	}
L15:
	;
	return
L16:
	;
	goto L14
L17:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[465]))
	v39 = *(*int32)(unsafe.Add(mBase, _consts[464]))
	v40 = v39
	v41 = v37
	goto L11
L18:
	;
	goto L8
L19:
	;
	F_MemoryContextDelete(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L15
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[467])) = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, _consts[468]))
	v67 = *(*int32)(unsafe.Add(mBase, _consts[469]))
	if v65|v67 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	if l0 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v97 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[466])) = uint8(v97)
	goto L3
L26:
	;
	v69 = int32(4487092)
	v70 = *(*int32)(unsafe.Add(mBase, _consts[258]))
	v73 = *(*int32)(unsafe.Add(mBase, _consts[402]))
	*(*int32)(unsafe.Add(mBase, _consts[258])) = v73
	if v67 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v89 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[469])) = v89
	*(*int32)(unsafe.Add(mBase, _consts[468])) = v89
	goto L25
L29:
	;
	F_relation_close(m, v67, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L15
	} else {
		goto L32
	}
L30:
	;
	v80 = v65
	goto L31
L31:
	;
	if v80 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _consts[468]))
	v80 = v79
	goto L31
L33:
	;
	F_sequence_close(m, v80, int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L15
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, _consts[258])) = v70
	goto L28
L36:
	;
	goto L35
}
