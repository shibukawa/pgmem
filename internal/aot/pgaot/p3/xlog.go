package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SendXlogRecPtrResult(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int64
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	v1 = l0
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+14)) = uint16(v3)
	v13 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_TupleDescInitBuiltinEntry(m, v16, int32(1), int32(_a_F_SendXlogRecPtrResult_0), int32(25))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_TupleDescInitBuiltinEntry(m, v16, int32(2), int32(_a_F_SendXlogRecPtrResult_1), int32(20))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v28 = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v28 < v37 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v116 = F_begin_tup_output_tupdesc(m, v13, v16, int32(_a_F_SendXlogRecPtrResult_2))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L25
	}
L7:
	;
	v41 = v16 + int32(28)
	v48 = v28
	v49 = v37
	v51 = v28
	goto L11
L8:
	;
	v105 = v28
	v112 = v37
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v105
	goto L6
L10:
	;
	v105 = v99
	v112 = v78
	goto L9
L11:
	;
	v57 = v41 + v37<<(uint(int32(3))%32) + v48*int32(100)
	v60 = v41 + v48<<(uint(int32(3))%32)
	if v37 != v49 {
		v78 = v49
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v99 = v37
	goto L10
L13:
	;
	v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+2)))
	if v79 <= int32(0) {
		v99 = v48
		goto L10
	} else {
		goto L21
	}
L14:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+7)))
	if v62 != int32(118) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v78 = v48
	goto L13
L16:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)))
	if v65 != int32(1) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+6)))
	if v68&int32(6) != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v71 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+2)))
	if v71 <= int32(0) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+90)))
	if v74 != int32(118) {
		v78 = v37
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+90)))
	if v82 == int32(118) {
		v99 = v48
		goto L10
	} else {
		goto L22
	}
L22:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+5)))
	v91 = (v51 + v85 - int32(1)) & (int32(0) - v85)
	if int32(_a_F_SendXlogRecPtrResult_3) < v91 {
		v99 = v48
		goto L10
	} else {
		goto L23
	}
L23:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60))) = uint16(v91)
	v97 = v48 + int32(1)
	if v97 != v37 {
		v48 = v97
		v49 = v78
		v51 = v91 + v79
		goto L11
	} else {
		goto L24
	}
L24:
	;
	goto L12
L25:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+4)) = uint32(v1)
	v120 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v8))) = uint32(v120)
	v123 = F_psprintf(m, int32(_a_F_SendXlogRecPtrResult_4), v8)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v125 = F_cstring_to_text(m, v123)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = base.I64_extend_i32_u(l1)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = base.I64_extend_i32_u(v125)
	F_do_tup_output(m, v116, v8+int32(16), v8+int32(14))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_end_tup_output(m, v116)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_pq_puttextmessage(m)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	m.G0 = v8 + int32(32)
	return
}
