package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AfterTriggerEndQuery(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	v9 = int32(4433368)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[401]))
	v13 = *(*int32)(unsafe.Add(mBase, _consts[402]))
	if v11 < v13 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[403]))
	v19 = v16 + v11*int32(20)
	v22 = F_afterTriggerMarkEvents(m, v19, int32(4433348), int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v148 = v11
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[401])) = v148 - int32(1)
	return
L4:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _consts[403]))
	v132 = *(*int32)(unsafe.Add(mBase, _consts[401]))
	F_AfterTriggerFreeQuery(m, v130+v132*int32(20))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
	} else {
		goto L30
	}
L5:
	;
	return
L6:
	;
	if v22 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v28 = v19
	goto L8
L8:
	;
	v34 = int32(4433340)
	v36 = *(*int32)(unsafe.Add(mBase, _consts[406]))
	*(*int32)(unsafe.Add(mBase, _consts[406])) = v36 + int32(1)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v42 = F_afterTriggerInvokeEvents(m, v28, v36, l0, int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L10
	}
L9:
	;
	goto L4
L10:
	;
	if v42 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[403]))
	v47 = *(*int32)(unsafe.Add(mBase, _consts[401]))
	v50 = v45 + v47*int32(20)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v40 != v51 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v56 = v51
	goto L15
L13:
	;
	goto L14
L14:
	;
	v119 = F_afterTriggerMarkEvents(m, v50, int32(4433348), int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L28
	}
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	if v61 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v103
	F_pfree(m, v56)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L26
	}
L18:
	;
	v64 = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v65 <= v64 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v69 = v64
	v75 = v65
	goto L20
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76+v69<<(uint(int32(2))%32))))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)))
	if v81 != int32(1) {
		v91 = v75
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L17
L22:
	;
	v93 = v69 + int32(1)
	if v93 < v91 {
		v69 = v93
		v75 = v91
		goto L20
	} else {
		goto L25
	}
L23:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	if v84 != v56 {
		v91 = v75
		goto L22
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v80)+12)) = int64(0)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v91 = v90
	goto L22
L25:
	;
	goto L21
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v107 != v40 {
		v56 = v107
		goto L15
	} else {
		goto L27
	}
L27:
	;
	goto L16
L28:
	;
	if v119 != 0 {
		v28 = v50
		goto L8
	} else {
		goto L29
	}
L29:
	;
	goto L9
L30:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[401]))
	v148 = v139
	goto L3
}
