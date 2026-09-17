package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BlockRefTableReaderGetBlocks(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	v10 = l0 + int32(_a_F_BlockRefTableReaderGetBlocks_0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[0])))
	v16 = v11
	v17 = int32(0)
	goto L1
L1:
	;
	if v16 == int32(0) {
		v108 = v17
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v108
L3:
	;
	if base.Ui32(l2) <= base.Ui32(v108) {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[1])))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[2])))
	v24 = int32(1)
	v25 = v16 - v24
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+v25<<(uint(v24)%32)))))
	if v29 != int32(_a_F_BlockRefTableReaderGetBlocks_1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if base.B2i32(base.Ui32(l2) <= base.Ui32(v17))|base.B2i32(base.Ui32(v29) <= base.Ui32(v22)) != 0 {
		v108 = v17
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if base.B2i32(base.Ui32(l2) <= base.Ui32(v17))|base.B2i32(base.Ui32(int32(_a_F_BlockRefTableReaderGetBlocks_2)) < base.Ui32(v22)) != 0 {
		v108 = v17
		goto L3
	} else {
		goto L13
	}
L8:
	;
	v40 = v22
	v42 = v17
	goto L9
L9:
	;
	v48 = int32(1)
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10+v40<<(uint(v48)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(l1+v42<<(uint(int32(2))%32)))) = v25<<(uint(int32(16))%32) | v51
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[1])))
	v56 = v54 + v48
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[1]))) = v56
	v59 = v42 + v48
	if base.Ui32(v29) <= base.Ui32(v56) {
		v108 = v59
		goto L3
	} else {
		goto L11
	}
L10:
	;
	v108 = v59
	goto L3
L11:
	;
	if base.Ui32(v59) < base.Ui32(l2) {
		v40 = v56
		v42 = v59
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v71 = v22
	v73 = v17
	goto L14
L14:
	;
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10+int32(base.Ui32(v71)>>(uint(int32(3))%32))&int32(536870910)))))
	if int32(base.Ui32(v81)>>(uint(v71&int32(15))%32))&int32(1) != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v108 = v95
	goto L3
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1+v73<<(uint(int32(2))%32)))) = v71 + v25<<(uint(int32(16))%32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[1])))
	v95 = v73 + int32(1)
	v96 = v94
	goto L18
L17:
	;
	v95 = v73
	v96 = v71
	goto L18
L18:
	;
	v98 = v96 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[1]))) = v98
	if base.Ui32(int32(_a_F_BlockRefTableReaderGetBlocks_2)) < base.Ui32(v98) {
		v108 = v95
		goto L3
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(v95) < base.Ui32(l2) {
		v71 = v98
		v73 = v95
		goto L14
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	goto L2
L22:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[0])))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[3])))
	if v112 == v113 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[2])))
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115+v112<<(uint(int32(1))%32)))))
	if v119 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_BlockRefTableRead(m, l0, v10, v119<<(uint(int32(1))%32))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v127 = v112
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[1]))) = int32(0)
	v131 = v127 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[0]))) = v131
	v16 = v131
	v17 = v108
	goto L1
L27:
	;
	return int32(0)
L28:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[0])))
	v127 = v126
	goto L26
}
func F_BlockStateAsString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	if base.Ui32(l0) <= base.Ui32(int32(19)) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_BlockStateAsString[0])))
		v8 = v6
	} else {
		v8 = int32(_a_F_BlockStateAsString_0)
	}
	return v8
}
func F_CreateBlockRefTableReader(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = F_palloc0(m, int32(_a_F_CreateBlockRefTableReader_0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_CreateBlockRefTableReader[0]))) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_CreateBlockRefTableReader[1]))) = int32(433)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_CreateBlockRefTableReader[2]))) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(432)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_CreateBlockRefTableReader[3]))) = int32(-1)
		F_BlockRefTableRead(m, v10, v7+int32(12), int32(4))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			if v29 != int32(1697321851) {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v29
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(1697321851)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_ReportWalSummaryError(m, int32(0), int32(_a_F_CreateBlockRefTableReader_1), v7)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(16)
					return v10
				}
			} else {
				m.G0 = v7 + int32(16)
				return v10
			}
		}
	}
}
func F_block_sampling_read_stream_next(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v36 float64
	_ = v36
	var v39 int32
	_ = v39
	var v41 float64
	_ = v41
	var v44 float64
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 float64
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 float64
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v4) < base.Ui32(v5) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v11 = base.B2i32(v7 < v8)
	} else {
		v11 = int32(0)
	}
	if v11 != 0 {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v21 = v19 - v20
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v24 = v22 - v23
		if base.Ui32(v21) < base.Ui32(v24) {
			for {
				v36 = F_pg_prng_double(m, l1+int32(16))
				mBase = m.M
				if base.F64_eq(v36, float64(0)) != 0 {
					continue
				} else {
					break
				}
				break
			}
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v41 = base.F64_convert_i32_s(v21)
			v44 = base.F64_sub(float64(1), base.F64_div(v41, base.F64_convert_i32_u(v24)))
			if base.F64_gt(v44, v36) != 0 {
				v47 = v39
				v49 = v24
				v52 = v44
				for {
					v54 = int32(1)
					v55 = v47 + v54
					v58 = v49 - v54
					v62 = base.F64_mul(v52, base.F64_sub(float64(1), base.F64_div(v41, base.F64_convert_i32_u(v58))))
					if base.F64_lt(v36, v62) != 0 {
						v47 = v55
						v49 = v58
						v52 = v62
						continue
					} else {
						break
					}
					break
				}
				v65 = v55
			} else {
				v65 = v39
			}
			v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v74 = v65
			v75 = v72
		} else {
			v74 = v23
			v75 = v20
		}
		v81 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v74 + v81
		*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v75 + v81
		v88 = v74
	} else {
		v88 = int32(-1)
	}
	return v88
}
