package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecCloseResultRelations(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v7 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v71 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v10 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = v2
	goto L4
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19+v17<<(uint(int32(2))%32))))
	F_ExecCloseIndices(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+212))
	if v26 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v62 = v17 + int32(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v62 < v63 {
		v17 = v62
		goto L4
	} else {
		goto L18
	}
L9:
	;
	v29 = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v30 <= v29 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v34 = v29
	goto L11
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v34<<(uint(int32(2))%32))))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v44 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L8
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	F_relation_close(m, v47, int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v52 = v34 + int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v52 < v53 {
		v34 = v52
		goto L11
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	goto L12
L18:
	;
	goto L5
L19:
	;
	return
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v74 <= int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v79 = int32(0)
	goto L22
L22:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84+v79<<(uint(int32(2))%32))))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	F_relation_close(m, v89, int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L24
	}
L23:
	;
	goto L19
L24:
	;
	v94 = v79 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v94 < v95 {
		v79 = v94
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
}
func F_remove_result_refs(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+80))
	if v14 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	v19 = F_get_relids_in_jointree(m, l2, int32(1), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v19
	v23 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l1
	v30 = F_query_or_expression_tree_walker_impl(m, v21, int32(900), v11+int32(4), v23)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v32 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v35 <= int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v45 = int32(-1)
	v46 = v4
	goto L8
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v46<<(uint(int32(2))%32))))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	if l1 == v52 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L1
L10:
	;
	if v45 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v60 = v45
	goto L12
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+80))
	if v62 != 0 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v56 = F_bms_singleton_member(m, v19)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L3
	} else {
		goto L16
	}
L14:
	;
	v58 = v45
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v58
	v60 = v58
	goto L12
L16:
	;
	v58 = v56
	goto L15
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v19
	v65 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l1
	v72 = F_query_or_expression_tree_walker_impl(m, v63, int32(900), v11+int32(4), v65)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v75 = v46 + int32(1)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v75 < v76 {
		v45 = v60
		v46 = v75
		goto L8
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	goto L9
}
