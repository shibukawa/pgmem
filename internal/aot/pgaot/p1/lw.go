package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LWLockDisownInternal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int64
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int64
	_ = v92
	var v96 int64
	_ = v96
	var v100 int64
	_ = v100
	var v103 int32
	_ = v103
	var v108 int64
	_ = v108
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[806]))
	v15 = v13
	v16 = int32(0)
	goto L2
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v22 = v15 - int32(1)
	if v22 < int32(0) {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v22<<(uint(int32(3))%32))+uint32(_consts[807])))
	v41 = v13 - int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[806])) = v41
	if v41 <= v22 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22<<(uint(int32(3))%32))+uint32(_consts[808])))
	if l0 != v31 {
		v15 = v22
		v16 = v16 + int32(1)
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	m.G0 = v10 + int32(16)
	return v37
L7:
	;
	v47 = v16 & int32(3)
	if v47 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v49 = v22
	v50 = int32(0)
	goto L11
L9:
	;
	v71 = v22
	goto L10
L10:
	;
	if base.Ui32(v16-int32(1)) < base.Ui32(int32(3)) {
		goto L6
	} else {
		goto L14
	}
L11:
	;
	v55 = int32(3)
	v59 = int32(1)
	v60 = v49 + v59
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v60<<(uint(v55)%32))+uint32(_consts[808])))
	*(*int64)(unsafe.Add(mBase, uint32(v49<<(uint(v55)%32))+uint32(_consts[808]))) = v65
	v68 = v50 + v59
	if v68 != v47 {
		v49 = v60
		v50 = v68
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v71 = v60
	goto L10
L13:
	;
	goto L12
L14:
	;
	v80 = v71
	goto L15
L15:
	;
	v86 = int32(3)
	v87 = v80 << (uint(v86) % 32)
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v87)+uint32(_consts[809])))
	*(*int64)(unsafe.Add(mBase, uint32(v87)+uint32(_consts[808]))) = v92
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v87)+uint32(_consts[810])))
	*(*int64)(unsafe.Add(mBase, uint32(v87)+uint32(_consts[809]))) = v96
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v87)+uint32(_consts[811])))
	*(*int64)(unsafe.Add(mBase, uint32(v87)+uint32(_consts[810]))) = v100
	v103 = v80 + int32(4)
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v103<<(uint(v86)%32))+uint32(_consts[808])))
	*(*int64)(unsafe.Add(mBase, uint32(v87)+uint32(_consts[811]))) = v108
	if v103 != v41 {
		v80 = v103
		goto L15
	} else {
		goto L17
	}
L16:
	;
	goto L6
L17:
	;
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v128) <= base.Ui32(int32(94)) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v155
	F_errmsg_internal(m, int32(430639), v10)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L18
	} else {
		goto L30
	}
L21:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v128<<(uint(int32(2))%32))+uint32(_consts[812])))
	v155 = v135
	goto L20
L22:
	;
	goto L23
L23:
	;
	v139 = (v128 - int32(95)) & int32(65535)
	v141 = *(*int32)(unsafe.Add(mBase, _consts[813]))
	if v139 < v141 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[814]))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v144+v139<<(uint(int32(2))%32))))
	if v148 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v153 = int32(270954)
	goto L26
L26:
	;
	v155 = v153
	goto L20
L27:
	;
	v150 = v148
	goto L29
L28:
	;
	v150 = int32(270954)
	goto L29
L29:
	;
	v153 = v150
	goto L26
L30:
	;
	F_errfinish(m, int32(495859), int32(1820), int32(311726))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L18
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_LWLockReleaseClearVar(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
	v5 = F_LWLockDisownInternal(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		F_LWLockReleaseInternal(m, l0, v5)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = int32(4483804)
			v11 = *(*int32)(unsafe.Add(mBase, _consts[163]))
			*(*int32)(unsafe.Add(mBase, _consts[163])) = v11 - int32(1)
			return
		}
	}
}
