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
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_ComboCid[0])) = v2
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_ComboCid[1])) = v2
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_ComboCid[2])) = v2
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_ComboCid[3])) = v2
	return
}
func F_AtEOXact_Inval(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Inval[0])) = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Inval[1]))
	if v9 != 0 {
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
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Inval[1])) = int32(0)
	goto L3
L5:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
	if v10 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if v55 < v56 {
		goto L23
	} else {
		goto L24
	}
L8:
	;
	F_RelationCacheInitFilePreInvalidate(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v17 = v9
	goto L10
L10:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v18
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v21
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v25 = v18 - v24
	if int32(0) < v25 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	return
L12:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Inval[1]))
	v17 = v16
	goto L10
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Inval[2]))
	F_SIInsertDataEntries(m, v29+v24<<(uint(int32(4))%32), v25)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	v36 = v21
	goto L15
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v38 = v36 - v37
	if int32(0) < v38 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v36 = v35
	goto L15
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Inval[3]))
	F_SIInsertDataEntries(m, v42+v37<<(uint(int32(4))%32), v38)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L11
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Inval[1]))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+16)))
	if v50 != int32(1) {
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
	v54 = m.ExcPending
	if v54 != 0 {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	goto L4
L23:
	;
	v58 = v55
	goto L26
L24:
	;
	goto L25
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	if v77 <= v76 {
		goto L4
	} else {
		goto L30
	}
L26:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Inval[2]))
	F_LocalExecuteInvalidationMessage(m, v63+v58<<(uint(int32(4))%32))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L11
	} else {
		goto L28
	}
L27:
	;
	goto L25
L28:
	;
	v70 = v58 + int32(1)
	if v70 != v56 {
		v58 = v70
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v79 = v76
	goto L31
L31:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Inval[3]))
	F_LocalExecuteInvalidationMessage(m, v84+v79<<(uint(int32(4))%32))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L11
	} else {
		goto L33
	}
L32:
	;
	goto L4
L33:
	;
	v91 = v79 + int32(1)
	if v91 != v77 {
		v79 = v91
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
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[0])))
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
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[1])) = v52
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[2])) = v52
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[3]))
	if v58 != 0 {
		goto L19
	} else {
		goto L20
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[1]))
	if v12 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[2]))
	v18 = v12
	v19 = v16
	v20 = int32(0)
	goto L7
L7:
	;
	v25 = v19 + v20<<(uint(int32(2))%32)
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
	v41 = v19
	goto L11
L11:
	;
	v43 = v20 + int32(1)
	if v43 < v40 {
		v18 = v40
		v19 = v41
		v20 = v43
		goto L7
	} else {
		goto L18
	}
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[4]))
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
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[2]))
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[1]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[3])) = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[5]))
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[6]))
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
	*(*uint8)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[0])) = uint8(v97)
	goto L3
L26:
	;
	v69 = int32(_a_F_AtEOXact_LargeObject_0)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[7]))
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[7])) = v73
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
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[6])) = v89
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[5])) = v89
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
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[5]))
	v80 = v79
	goto L31
L33:
	;
	F_relation_close(m, v80, int32(0))
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
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LargeObject[7])) = v70
	goto L28
L36:
	;
	goto L35
}
