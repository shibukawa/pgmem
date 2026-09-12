package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AcquireExecutorLocks(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = v3
	goto L4
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v17<<(uint(int32(2))%32))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v25 == int32(6) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L1
L6:
	;
	v102 = v17 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v102 < v103 {
		v17 = v102
		goto L4
	} else {
		goto L41
	}
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+88))
	v30 = v28
	goto L11
L8:
	;
	goto L9
L9:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v59 == int32(0) {
		goto L6
	} else {
		goto L27
	}
L10:
	;
	if v54 == int32(0) {
		goto L6
	} else {
		goto L24
	}
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	switch v32 - int32(241) {
	case 0:
		goto L16
	case 1:
		goto L15
	default:
		goto L17
	}
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+28))
	v30 = v52
	goto L11
L14:
	;
	v54 = v50
	goto L10
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v47 == int32(6) {
		v51 = v46
		goto L13
	} else {
		goto L23
	}
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v43 == int32(6) {
		v51 = v42
		goto L13
	} else {
		goto L22
	}
L17:
	;
	if v32 != int32(201) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v54 = int32(0)
	goto L10
L19:
	;
	goto L20
L20:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v39 != int32(6) {
		v50 = v38
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v51 = v38
	goto L13
L22:
	;
	v50 = v42
	goto L14
L23:
	;
	v50 = v46
	goto L14
L24:
	;
	F_ScanQueryForLocks(m, v54, l1)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return
L26:
	;
	goto L6
L27:
	;
	v62 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v63 <= v62 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v68 = v62
	goto L29
L29:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v68<<(uint(int32(2))%32))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	switch v78 {
	case 0:
		goto L32
	case 1:
		goto L33
	default:
		goto L31
	}
L30:
	;
	goto L6
L31:
	;
	v91 = v68 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v91 < v92 {
		v68 = v91
		goto L29
	} else {
		goto L40
	}
L32:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)+24))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	if l1 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	if v79 == int32(0) {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	F_LockRelationOid(m, v83, v82)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L25
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	F_UnlockRelationOid(m, v83, v82)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L25
	} else {
		goto L39
	}
L38:
	;
	goto L31
L39:
	;
	goto L31
L40:
	;
	goto L30
L41:
	;
	goto L5
}
func F_ExecutorStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int64
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)+8))
	v9 = *(*int32)(unsafe.Add(mBase, _consts[49]))
	if v9 == int32(0) {
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, _consts[50])))
		if v13 != int32(1) {
		} else {
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v9)+392))
			if int32(1)&base.B2i32(v18 != int64(0)) != 0 {
			} else {
				v22 = int32(4556756)
				v24 = *(*int32)(unsafe.Add(mBase, _consts[14]))
				v25 = int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[14])) = v24 + v25
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v28 + v25
				*(*int64)(unsafe.Add(mBase, uint32(v9)+392)) = v5
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v28 + int32(2)
				v39 = *(*int32)(unsafe.Add(mBase, _consts[14]))
				*(*int32)(unsafe.Add(mBase, _consts[14])) = v39 - v25
			}
		}
	}
	v44 = *(*int32)(unsafe.Add(mBase, _consts[336]))
	if v44 != 0 {
		m.T0[v44].(func(*base.Module, int32, int32))(m, l0, l1)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return
		} else {
			return
		}
	} else {
		F_standard_ExecutorStart(m, l0, l1)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			return
		}
	}
}
