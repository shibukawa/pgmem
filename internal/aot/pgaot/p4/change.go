package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_changeDependenciesOf(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	v8 = m.G0
	v10 = v8 - int32(112)
	m.G0 = v10
	v14 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_ScanKeyInit(m, v10, int32(1), int32(3), int32(184), int64(1259))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v10+int32(56), int32(2), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = F_systable_beginscan(m, v14, int32(2673), int32(1), int32(0), int32(2), v10)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v37 = F_systable_getnext(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v37 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v41 = v37
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_systable_endscan(m, v35)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L17
	}
L10:
	;
	v46 = F_heap_copytuple(m, v41)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v48+v49)+4)) = l1
	F_CatalogTupleUpdate(m, v14, v46+int32(4), v46)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_pfree(m, v46)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v60 = F_systable_getnext(m, v35)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v60 != 0 {
		v41 = v60
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L11
L17:
	;
	F_relation_close(m, v14, int32(3))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	m.G0 = v10 + int32(112)
	return
}
func F_change_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(_a_F_change_cb_wrapper_0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v12
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(1058)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v16
	v20 = int32(_a_F_change_cb_wrapper_1)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_change_cb_wrapper[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_change_cb_wrapper[0])) = v10 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v10 + int32(16)
	v30 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+147)) = uint8(v30)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v32
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+164)) = uint8(v5)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+152)) = v34
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	m.T0[v38].(func(*base.Module, int32, int32, int32, int32))(m, v12, l1, l2, l3)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		return
	} else {
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		*(*int32)(unsafe.Add(mBase, _c_F_change_cb_wrapper[0])) = v42
		m.G0 = v10 + int32(32)
		return
	}
}
