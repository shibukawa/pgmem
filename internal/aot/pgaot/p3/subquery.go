package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_set_subquery_size_estimates(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v96 int32
	_ = v96
	v3 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v14 = F_fetch_upper_rel(m, v11, int32(7), v3)
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v17 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+120)) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v20 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_set_baserel_size_estimates(m, l0, l1)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L17
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v23 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v29 = v3
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v29<<(uint(int32(2))%32))))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+26)))
	if v41 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L3
L8:
	;
	v82 = v29 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v82 < v83 {
		v29 = v82
		goto L6
	} else {
		goto L16
	}
L9:
	;
	v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v40)+8)))
	v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+80)))
	if v42 < v43 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+82)))
	if v45 < v42 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v47 = int32(0)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v49 != int32(6) {
		v67 = v42
		v68 = v43
		v69 = v47
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v70+(base.I32_extend16_s(v67)-v68)<<(uint(int32(2))%32)))) = v69
	goto L8
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+144))
	if v53 != 0 {
		v67 = v42
		v68 = v43
		v69 = v47
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v55 = F_find_base_rel(m, v11, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+88))
	v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+8)))
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55)+80)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v57+(v58-v59)<<(uint(int32(2))%32))))
	v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+80)))
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+8)))
	v67 = v66
	v68 = v65
	v69 = v64
	goto L12
L16:
	;
	goto L7
L17:
	;
	return
}
