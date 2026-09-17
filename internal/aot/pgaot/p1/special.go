package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DecodeSpecial(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	v3 = int32(_a_F_DecodeSpecial_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeSpecial[0]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v78
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeSpecial[0])) = v59
	v65 = int32(*(*int8)(unsafe.Add(mBase, uint32(v59)+11)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v66
	v78 = v65
	goto L1
L3:
	;
	v13 = F_strncmp(m, l0, v11, int32(10))
	mBase = m.M
	if v13 == int32(0) {
		v59 = v11
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	v23 = int32(_a_F_DecodeSpecial_1)
	v25 = int32(_a_F_DecodeSpecial_2)
	goto L7
L6:
	;
	goto L5
L7:
	;
	v32 = v23 + (v25-v23)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32))))
	v34 = v16 - v33
	if v34 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v78 = int32(31)
	goto L1
L9:
	;
	v38 = F_strncmp(m, l0, v32, int32(10))
	mBase = m.M
	if v38 == int32(0) {
		v59 = v32
		goto L2
	} else {
		goto L12
	}
L10:
	;
	v41 = v34
	goto L11
L11:
	;
	v45 = base.B2i32(v41 < int32(0))
	if v41 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v41 = v38
	goto L11
L13:
	;
	v46 = v32 - int32(16)
	goto L15
L14:
	;
	v46 = v25
	goto L15
L15:
	;
	if v41 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v49 = v23
	goto L18
L17:
	;
	v49 = v32 + int32(16)
	goto L18
L18:
	;
	if base.Ui32(v49) <= base.Ui32(v46) {
		v23 = v49
		v25 = v46
		goto L7
	} else {
		goto L19
	}
L19:
	;
	goto L8
}
func F_EncodeSpecialDate(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	if l0 != int32(2147483647) {
		if l0 == int32(-2147483648) {
			v8 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_EncodeSpecialDate[0])))
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v8)
			v11 = *(*int64)(unsafe.Add(mBase, _c_F_EncodeSpecialDate[1]))
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v11
			return
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_EncodeSpecialDate_0), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_EncodeSpecialDate_1), int32(309), int32(_a_F_EncodeSpecialDate_2))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EncodeSpecialDate[2])))
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v27)
		v30 = *(*int64)(unsafe.Add(mBase, _c_F_EncodeSpecialDate[3]))
		*(*int64)(unsafe.Add(mBase, uint32(l1))) = v30
		return
	}
}
