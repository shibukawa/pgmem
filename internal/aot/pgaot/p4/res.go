package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResOwnerPrintBuffer(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v67 int64
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(112)
	m.G0 = v10
	v12 = base.I32_wrap_i64(l0)
	if v12 < v2 {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBuffer[0]))
		v18 = v12 ^ int32(-1)
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBuffer[1]))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v18<<(uint(int32(2))%32))))
		v29 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBuffer[2]))
		v60 = v16 + v18*int32(56)
		v62 = v27
		v63 = v29
		v64 = int64(0)
		v67 = base.AtomicRmwCmpxchg64(m, v60, int32(24), v64, v64)
		v69 = v10 + int32(40)
		v70 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
		v71 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
		v72 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
		v73 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
		F_GetRelationPath(m, v69, v70, v71, v72, v63, v73)
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			v76 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v62
			*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = base.I32_wrap_i64(v67) & int32(_a_F_ResOwnerPrintBuffer_0)
			*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v67 & int64(17175674880)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v76
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v12
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v69
			v89 = F_psprintf(m, int32(_a_F_ResOwnerPrintBuffer_1), v10)
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(112)
				return v89
			}
		}
	} else {
		v31 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBuffer[3]))
		v32 = int32(56)
		v38 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerPrintBuffer[4]))
		if v38 != int32(-1) {
			v42 = v38 << (uint(int32(4)) % 32)
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_ResOwnerPrintBuffer[5])))
			if v45 == v12 {
				v55 = v42 + int32(_a_F_ResOwnerPrintBuffer_2)
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
				v58 = v56
				v60 = v31 + v12*v32 - v32
				v62 = v58
				v63 = int32(-1)
				v64 = int64(0)
				v67 = base.AtomicRmwCmpxchg64(m, v60, int32(24), v64, v64)
				v69 = v10 + int32(40)
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
				v72 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
				F_GetRelationPath(m, v69, v70, v71, v72, v63, v73)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v62
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = base.I32_wrap_i64(v67) & int32(_a_F_ResOwnerPrintBuffer_0)
					*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v67 & int64(17175674880)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v76
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v12
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v69
					v89 = F_psprintf(m, int32(_a_F_ResOwnerPrintBuffer_1), v10)
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 + int32(112)
						return v89
					}
				}
			} else {
				v49 = F_GetPrivateRefCountEntrySlow(m, v12, int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					if v49 == int32(0) {
						v58 = v2
					} else {
						v55 = v49
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
						v58 = v56
					}
					v60 = v31 + v12*v32 - v32
					v62 = v58
					v63 = int32(-1)
					v64 = int64(0)
					v67 = base.AtomicRmwCmpxchg64(m, v60, int32(24), v64, v64)
					v69 = v10 + int32(40)
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
					F_GetRelationPath(m, v69, v70, v71, v72, v63, v73)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v62
						*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = base.I32_wrap_i64(v67) & int32(_a_F_ResOwnerPrintBuffer_0)
						*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v67 & int64(17175674880)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v76
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v12
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v69
						v89 = F_psprintf(m, int32(_a_F_ResOwnerPrintBuffer_1), v10)
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							m.G0 = v10 + int32(112)
							return v89
						}
					}
				}
			}
		} else {
			v49 = F_GetPrivateRefCountEntrySlow(m, v12, int32(0))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				if v49 == int32(0) {
					v58 = v2
				} else {
					v55 = v49
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
					v58 = v56
				}
				v60 = v31 + v12*v32 - v32
				v62 = v58
				v63 = int32(-1)
				v64 = int64(0)
				v67 = base.AtomicRmwCmpxchg64(m, v60, int32(24), v64, v64)
				v69 = v10 + int32(40)
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
				v72 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
				F_GetRelationPath(m, v69, v70, v71, v72, v63, v73)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v62
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = base.I32_wrap_i64(v67) & int32(_a_F_ResOwnerPrintBuffer_0)
					*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v67 & int64(17175674880)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v76
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v12
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v69
					v89 = F_psprintf(m, int32(_a_F_ResOwnerPrintBuffer_1), v10)
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 + int32(112)
						return v89
					}
				}
			}
		}
	}
}
func F_ResOwnerPrintCatCache(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = base.I32_wrap_i64(l0)
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+6)))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+84))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v14 | v15<<(uint(int32(16))%32)
	v32 = F_psprintf(m, int32(_a_F_ResOwnerPrintCatCache_0), v11)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return int32(0)
	} else {
		m.G0 = v11 + int32(32)
		return v32
	}
}
func F_ResOwnerPrintDSM(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))+12))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v9
	v12 = F_psprintf(m, int32(_a_F_ResOwnerPrintDSM_0), v6)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v12
	}
}
func F_ResOwnerReleaseBufferIO(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v26 int64
	_ = v26
	var v40 int64
	_ = v40
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v59 int64
	_ = v59
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v100 int64
	_ = v100
	var v107 int32
	_ = v107
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBufferIO[0]))
	v16 = v12 + base.I32_wrap_i64(l0)*int32(56)
	v18 = v16 - int32(32)
	v19 = int64(4194304)
	v21 = base.AtomicRmwOr64(m, v18, int32(0), v19)
	if v21&v19 != int64(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = v21
	goto L4
L2:
	;
	v100 = v21
	goto L3
L3:
	;
	v107 = v16 - int32(56)
	v110 = base.AtomicRmwSub64(m, v18, int32(0), int64(4194304))
	v111 = int64(150994944)
	if v100&v111 != v111 {
		goto L26
	} else {
		goto L27
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = int32(_a_F_ResOwnerReleaseBufferIO_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a_F_ResOwnerReleaseBufferIO_1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_ResOwnerReleaseBufferIO_2)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(0)
	v40 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v40
	if v26&int64(4194304) != v40 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v100 = v95
	goto L3
L6:
	;
	goto L9
L7:
	;
	goto L8
L8:
	;
	v73 = int32(_a_F_ResOwnerReleaseBufferIO_3)
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBufferIO[1]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(8))+8))
	if v76 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	F_perform_spin_delay(m, v9+int32(8))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	return
L12:
	;
	v56 = int64(0)
	v59 = base.AtomicRmwCmpxchg64(m, v18, int32(0), v56, v56)
	if v59&int64(4194304) != v56 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v93 = int64(4194304)
	v95 = base.AtomicRmwOr64(m, v18, int32(0), v93)
	if v95&v93 != int64(0) {
		v26 = v95
		goto L4
	} else {
		goto L25
	}
L15:
	;
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBufferIO[1])) = v91
	goto L15
L17:
	;
	if int32(999) < v74 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v74 < int32(11) {
		goto L15
	} else {
		goto L24
	}
L20:
	;
	v81 = int32(900)
	if v81 <= v74 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v84 = v81
	goto L23
L22:
	;
	v84 = v74
	goto L23
L23:
	;
	v91 = v84 + int32(100)
	goto L16
L24:
	;
	v91 = v74 - int32(1)
	goto L16
L25:
	;
	goto L5
L26:
	;
	v158 = int32(0)
	F_TerminateBufferIO(m, v107, v158, int64(134217728), v158, v158)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L11
	} else {
		goto L35
	}
L27:
	;
	v117 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	if v117 == int32(0) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	F_errcode(m, int32(_a_F_ResOwnerReleaseBufferIO_4))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v16-int32(40))))
	v128 = v9 + int32(8)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v16-int32(52))))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v16-int32(48))))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v16-int32(44))))
	F_GetRelationPath(m, v128, v131, v132, v135, int32(-1), v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L11
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v128
	F_errmsg(m, int32(_a_F_ResOwnerReleaseBufferIO_5), v9)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L11
	} else {
		goto L32
	}
L32:
	;
	v149 = F_errdetail(m, int32(_a_F_ResOwnerReleaseBufferIO_6), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L11
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_ResOwnerReleaseBufferIO_2), int32(_a_F_ResOwnerReleaseBufferIO_7), int32(_a_F_ResOwnerReleaseBufferIO_8))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	goto L26
L35:
	;
	m.G0 = v9 + int32(80)
	return
}
func F_ResOwnerReleaseCatCacheList(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = base.I32_wrap_i64(l0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+48))
	v6 = int32(1)
	v7 = v5 - v6
	*(*int32)(unsafe.Add(mBase, uint32(v4)+48)) = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+52)))
	if base.B2i32(v9 != v6)|v7 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)+60))
		F_CatCacheRemoveCList(m, v15, v4)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_ResOwnerReleaseDSM(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v3 = base.I32_wrap_i64(l0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = int32(0)
	F_dsm_detach(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_ResOwnerReleaseFile(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseFile[0]))
	v5 = base.I32_wrap_i64(l0)
	*(*int32)(unsafe.Add(mBase, uint32(v4+v5*int32(48))+8)) = int32(0)
	F_FileClose(m, v5)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		return
	}
}
func F_ResOwnerReleaseSnapshot(m *base.Module, l0 int64) {
	var v4 int32
	_ = v4
	F_UnregisterSnapshotNoOwner(m, base.I32_wrap_i64(l0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
