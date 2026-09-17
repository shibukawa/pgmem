package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pq_beginmessage_reuse(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
	return
}
func F_pq_getbytes(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v9 = l0
	v10 = l1
	goto L4
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_pq_getbytes[0]))
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_pq_getbytes[1]))
	if v14 <= v16 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v18 = F_pq_recvbuf(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v26 = v14 - v16
	if base.Ui32(v26) < base.Ui32(v10) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	return int32(0)
L10:
	;
	if v18 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	return int32(-1)
L12:
	;
	v28 = v26
	goto L14
L13:
	;
	v28 = v10
	goto L14
L14:
	;
	if v28 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	base.MemoryCopy(m, v9, v16+int32(_a_F_pq_getbytes_0), v28)
	goto L17
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pq_getbytes[1])) = v28 + v16
	v36 = v10 - v28
	if v36 != 0 {
		v9 = v9 + v28
		v10 = v36
		goto L4
	} else {
		goto L18
	}
L18:
	;
	goto L5
}
func F_pq_getmsgfloat4(m *base.Module, l0 int32) float32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v4-v5 <= int32(3) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return float32(0)
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return float32(0)
			} else {
				F_errmsg(m, int32(_a_F_pq_getmsgfloat4_0), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return float32(0)
				} else {
					F_errfinish(m, int32(_a_F_pq_getmsgfloat4_1), int32(533), int32(_a_F_pq_getmsgfloat4_2))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return float32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v27+v5)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5 + int32(4)
		v35 = int32(16711935)
		return base.F32_reinterpret_i32(base.I32_rotr(v29, int32(24))&v35 | base.I32_rotr(v29&v35, int32(8)))
	}
}
func F_pq_getmsgfloat8(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v4-v5 <= int32(7) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return float64(0)
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return float64(0)
			} else {
				F_errmsg(m, int32(_a_F_pq_getmsgfloat8_0), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return float64(0)
				} else {
					F_errfinish(m, int32(_a_F_pq_getmsgfloat8_1), int32(533), int32(_a_F_pq_getmsgfloat8_2))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return float64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v29 = *(*int64)(unsafe.Add(mBase, uint32(v27+v5)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5 + int32(8)
		v33 = int64(56)
		v35 = int64(65280)
		v37 = int64(40)
		v40 = int64(16711680)
		v42 = int64(24)
		v44 = int64(4278190080)
		v46 = int64(8)
		return base.F64_reinterpret_i64(v29<<(uint(v33)%64) | v29&v35<<(uint(v37)%64) | (v29&v40<<(uint(v42)%64) | v29&v44<<(uint(v46)%64)) | (int64(base.Ui64(v29)>>(uint(v46)%64))&v44 | int64(base.Ui64(v29)>>(uint(v42)%64))&v40 | (int64(base.Ui64(v29)>>(uint(v37)%64))&v35 | int64(base.Ui64(v29)>>(uint(v33)%64)))))
	}
}
