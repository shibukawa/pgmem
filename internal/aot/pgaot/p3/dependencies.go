package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dependencies_array_element_start(m *base.Module, l0 int32, l1 int32) int32 {
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v10 = Fn14253(m, l0, l1, int32(_a_F_dependencies_array_element_start_0), int32(441), int32(_a_F_dependencies_array_element_start_1), int32(_a_F_dependencies_array_element_start_2), int32(435), int32(_a_F_dependencies_array_element_start_3), int32(425))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_dependencies_array_end(m *base.Module, l0 int32) int32 {
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v10 = Fn14254(m, l0, int32(_a_F_dependencies_array_end_0), int32(325), int32(_a_F_dependencies_array_end_1), int32(_a_F_dependencies_array_end_2), int32(314), int32(_a_F_dependencies_array_end_3), int32(7), int32(301))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_generate_dependencies_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9-int32(1) <= l1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v91 <= l2 {
		goto L1
	} else {
		goto L22
	}
L5:
	;
	v23 = int32(0)
	goto L6
L6:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l3+l1<<(uint(int32(1))%32)))) = uint16(v23)
	v28 = int32(0)
	if v28 < l1 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L1
L8:
	;
	v88 = v23 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v88 < v89 {
		v23 = v88
		goto L6
	} else {
		goto L21
	}
L9:
	;
	v33 = v28
	goto L12
L10:
	;
	goto L11
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v57 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	v58 = int32(1)
	v63 = F_repalloc(m, v55, v56*(v57+v58)<<(uint(v58)%32))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(l3+v33<<(uint(int32(1))%32)))))
	if v23 == v42 {
		goto L8
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	v45 = v33 + int32(1)
	if v45 != l1 {
		v33 = v45
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	return
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v63
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v68 = v66 << (uint(int32(1)) % 32)
	if v68 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	base.MemoryCopy(m, v63+v66*v69<<(uint(int32(1))%32), l3, v68)
	goto L20
L19:
	;
	goto L20
L20:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v77 = v75 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v77)
	goto L8
L21:
	;
	goto L7
L22:
	;
	v93 = int32(1)
	v100 = l2
	goto L23
L23:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l3+l1<<(uint(v93)%32)))) = uint16(v100)
	v109 = base.I32_extend16_s(v100 + int32(1))
	F_generate_dependencies_recurse(m, l0, l1+v93, v109, l3)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L16
	} else {
		goto L25
	}
L24:
	;
	goto L1
L25:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v109 < v112 {
		v100 = v109
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
}
