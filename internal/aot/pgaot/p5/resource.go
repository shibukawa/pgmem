package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResourceOwnerReleaseAll(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return
L2:
	;
	v86 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v86)
	goto L1
L3:
	;
	v26 = v21
	goto L9
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v21 = v14
	v22 = v15
	goto L3
L5:
	;
	goto L6
L6:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	if v16 == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v21 = v16
	v22 = l0 + int32(24)
	goto L3
L8:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v81 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L9:
	;
	v34 = v22 + v26<<(uint(int32(3))%32)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34-int32(4))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if base.Ui32(l1) < base.Ui32(v38) {
		v78 = v26
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v78 = int32(0)
	goto L8
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v34-int32(8))))
	if l2 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	if v43 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	m.T0[v72].(func(*base.Module, int32))(m, v42)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L19
	} else {
		goto L29
	}
L15:
	;
	v58 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L19
	} else {
		goto L22
	}
L16:
	;
	v44 = m.T0[v43].(func(*base.Module, int32) int32)(m, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v46
	v52 = F_psprintf(m, int32(239639), v12+int32(16))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L19
	} else {
		goto L21
	}
L19:
	;
	return
L20:
	;
	v55 = v44
	goto L15
L21:
	;
	v55 = v52
	goto L15
L22:
	;
	if v58 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v55
	F_errmsg_internal(m, int32(204398), v12)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L19
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	F_pfree(m, v55)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L19
	} else {
		goto L28
	}
L26:
	;
	F_errfinish(m, int32(496474), int32(395), int32(306122))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	goto L14
L29:
	;
	v76 = v26 - int32(1)
	if v76 != 0 {
		v26 = v76
		goto L9
	} else {
		goto L30
	}
L30:
	;
	goto L10
L31:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v78)
	goto L1
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v78
	goto L1
}
