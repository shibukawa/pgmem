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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	v10 = l0 + int32(_a_F_BlockRefTableReaderGetBlocks_0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[0])))
	v16 = int32(0)
	v17 = v11
	goto L1
L1:
	;
	if v17 == int32(0) {
		v105 = v16
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v105
L3:
	;
	if base.Ui32(l2) <= base.Ui32(v105) {
		goto L23
	} else {
		goto L24
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[1])))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[2])))
	v24 = int32(1)
	v25 = v17 - v24
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+v25<<(uint(v24)%32)))))
	if v29 != int32(_a_F_BlockRefTableReaderGetBlocks_1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if base.Ui32(v29) <= base.Ui32(v22) {
		v105 = v16
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if base.Ui32(int32(_a_F_BlockRefTableReaderGetBlocks_2)) < base.Ui32(v22) {
		v105 = v16
		goto L3
	} else {
		goto L14
	}
L8:
	;
	if base.Ui32(l2) <= base.Ui32(v16) {
		v105 = v16
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v39 = v22
	v40 = v16
	goto L10
L10:
	;
	v47 = int32(1)
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10+v39<<(uint(v47)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(l1+v40<<(uint(int32(2))%32)))) = v25<<(uint(int32(16))%32) | v50
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[1])))
	v55 = v53 + v47
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[1]))) = v55
	v58 = v40 + v47
	if base.Ui32(v29) <= base.Ui32(v55) {
		v105 = v58
		goto L3
	} else {
		goto L12
	}
L11:
	;
	v105 = v58
	goto L3
L12:
	;
	if base.Ui32(v58) < base.Ui32(l2) {
		v39 = v55
		v40 = v58
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	if base.Ui32(l2) <= base.Ui32(v16) {
		v105 = v16
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v69 = v22
	v70 = v16
	goto L16
L16:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10+int32(base.Ui32(v69)>>(uint(int32(3))%32))&int32(536870910)))))
	if int32(base.Ui32(v79)>>(uint(v69&int32(15))%32))&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v105 = v93
	goto L3
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1+v70<<(uint(int32(2))%32)))) = v69 + v25<<(uint(int32(16))%32)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[1])))
	v93 = v70 + int32(1)
	v94 = v92
	goto L20
L19:
	;
	v93 = v70
	v94 = v69
	goto L20
L20:
	;
	v96 = v94 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[1]))) = v96
	if base.Ui32(int32(_a_F_BlockRefTableReaderGetBlocks_2)) < base.Ui32(v96) {
		v105 = v93
		goto L3
	} else {
		goto L21
	}
L21:
	;
	if base.Ui32(v93) < base.Ui32(l2) {
		v69 = v96
		v70 = v93
		goto L16
	} else {
		goto L22
	}
L22:
	;
	goto L17
L23:
	;
	goto L2
L24:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[0])))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[3])))
	if v110 == v111 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[2])))
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113+v110<<(uint(int32(1))%32)))))
	if v117 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_BlockRefTableRead(m, l0, v10, v117<<(uint(int32(1))%32))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v125 = v110
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[1]))) = int32(0)
	v129 = v125 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[0]))) = v129
	v16 = v105
	v17 = v129
	goto L1
L29:
	;
	return int32(0)
L30:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderGetBlocks[0])))
	v125 = v124
	goto L28
}
func F_BlockStateAsString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	if base.Ui32(l0) <= base.Ui32(int32(19)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_BlockStateAsString[0])))
		v11 = v10
	} else {
		v11 = int32(_a_F_BlockStateAsString_0)
	}
	return v11
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
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v25 int32
	_ = v25
	var v37 float64
	_ = v37
	var v40 int32
	_ = v40
	var v42 float64
	_ = v42
	var v45 float64
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 float64
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 float64
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v6) < base.Ui32(v7) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v12 = base.B2i32(v9 < v10)
	} else {
		v12 = int32(0)
	}
	if v12 != 0 {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v22 = v20 - v21
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v25 = v23 - v24
		if base.Ui32(v22) < base.Ui32(v25) {
			for {
				v37 = F_pg_prng_double(m, l1+int32(16))
				mBase = m.M
				if base.F64_eq(v37, float64(0)) != 0 {
					continue
				} else {
					break
				}
				break
			}
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v42 = base.F64_convert_i32_s(v22)
			v45 = base.F64_sub(float64(1), base.F64_div(v42, base.F64_convert_i32_u(v25)))
			if base.F64_gt(v45, v37) != 0 {
				v48 = v40
				v50 = v25
				v53 = v45
				for {
					v55 = int32(1)
					v56 = v48 + v55
					v59 = v50 - v55
					v63 = base.F64_mul(v53, base.F64_sub(float64(1), base.F64_div(v42, base.F64_convert_i32_u(v59))))
					if base.F64_lt(v37, v63) != 0 {
						v48 = v56
						v50 = v59
						v53 = v63
						continue
					} else {
						break
					}
					break
				}
				v66 = v56
			} else {
				v66 = v40
			}
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v75 = v66
			v76 = v73
		} else {
			v75 = v24
			v76 = v21
		}
		v82 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v75 + v82
		*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v76 + v82
		v88 = v75
	} else {
		v88 = int32(-1)
	}
	return v88
}
