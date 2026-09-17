package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecEndSeqScan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v2 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
		v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+188))
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
		m.T0[v5].(func(*base.Module, int32))(m, v2)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_ExecSeqScanWithProject(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
	F_MemoryContextReset(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithProject[0]))
		if v5 == int32(0) {
			if v13 != 0 {
				F_ProcessInterrupts(m)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v18 = F_SeqNext(m, l0)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						return v18
					}
				}
			} else {
				v18 = F_SeqNext(m, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					return v18
				}
			}
		} else {
			if v13 != 0 {
				F_ProcessInterrupts(m)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = F_SeqNext(m, l0)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						if v23 != 0 {
							v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
							if v25&int32(2) == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v23
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v5)+72))
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
								m.T0[v40].(func(*base.Module, int32))(m, v38)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v43 = int32(_a_F_ExecSeqScanWithProject_0)
									v44 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithProject[1]))
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
									*(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithProject[1])) = v46
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v5)+24))
									v52 = m.T0[v51].(func(*base.Module, int32, int32, int32) int32)(m, v5+int32(4), v37, int32(0))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithProject[1])) = v44
										v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+4)))
										v58 = v56 & int32(_a_F_ExecSeqScanWithProject_1)
										*(*uint16)(unsafe.Add(mBase, uint32(v38)+4)) = uint16(v58)
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
										*(*uint16)(unsafe.Add(mBase, uint32(v38)+6)) = uint16(v61)
										return v38
									}
								}
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
								v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
								v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
								m.T0[v32].(func(*base.Module, int32))(m, v30)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int32(0)
								} else {
									return v30
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
							m.T0[v32].(func(*base.Module, int32))(m, v30)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								return v30
							}
						}
					}
				}
			} else {
				v23 = F_SeqNext(m, l0)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if v23 != 0 {
						v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
						if v25&int32(2) == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v23
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v5)+72))
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
							m.T0[v40].(func(*base.Module, int32))(m, v38)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = int32(_a_F_ExecSeqScanWithProject_0)
								v44 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithProject[1]))
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
								*(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithProject[1])) = v46
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v5)+24))
								v52 = m.T0[v51].(func(*base.Module, int32, int32, int32) int32)(m, v5+int32(4), v37, int32(0))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithProject[1])) = v44
									v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+4)))
									v58 = v56 & int32(_a_F_ExecSeqScanWithProject_1)
									*(*uint16)(unsafe.Add(mBase, uint32(v38)+4)) = uint16(v58)
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
									*(*uint16)(unsafe.Add(mBase, uint32(v38)+6)) = uint16(v61)
									return v38
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
							m.T0[v32].(func(*base.Module, int32))(m, v30)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								return v30
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
						m.T0[v32].(func(*base.Module, int32))(m, v30)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							return v30
						}
					}
				}
			}
		}
	}
}
func F_fill_seq_fork_with_data(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int64
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = l0
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v10)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v17
	v23 = F_ExtendBufferedRel(m, v10+int32(8), l2, int32(0), int32(9))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v43 = int32(_a_F_fill_seq_fork_with_data_0)
	v45 = int32(0)
	if v45|(v42&int32(3)|int32(1)) == v45 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	return
L3:
	;
	if v23 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_fill_seq_fork_with_data[0]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28+(v23^int32(-1))<<(uint(int32(2))%32))))
	v42 = v34
	goto L1
L5:
	;
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_fill_seq_fork_with_data[1]))
	v42 = v36 + v23<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L7:
	;
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v42+v93))) = int32(_a_F_fill_seq_fork_with_data_4)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = int32(2)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+20)))
	v103 = v101 | int32(768)
	*(*uint16)(unsafe.Add(mBase, uint32(v100)+20)) = uint16(v103)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v106 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v106
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+20)))
	v110 = v108 & int32(_a_F_fill_seq_fork_with_data_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v105)+20)) = uint16(v110)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+4)) = v106
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115)+20)))
	v118 = v116 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v115)+20)) = uint16(v118)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v121 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+16)) = uint16(v121)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+12)) = v106
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+118)))
	if v126 != int32(112) {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+10)) = int32(_a_F_fill_seq_fork_with_data_1)
	v84 = int32(_a_F_fill_seq_fork_with_data_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v42)+18)) = uint16(v84)
	v90 = int32(_a_F_fill_seq_fork_with_data_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v42)+16)) = uint16(v90)
	*(*uint16)(unsafe.Add(mBase, uint32(v42)+14)) = uint16(v90)
	goto L7
L9:
	;
	goto L12
L10:
	;
	goto L11
L11:
	;
	goto L17
L12:
	;
	v61 = v42 + v43
	v63 = v42 + int32(4)
	if base.Ui32(v63) < base.Ui32(v61) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v65 = v61
	goto L15
L14:
	;
	v65 = v63
	goto L15
L15:
	;
	v70 = (v42^int32(-1)+v65)&int32(-4) + int32(4)
	if v70 == int32(0) {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	base.MemoryFill(m, v42, int32(0), v70)
	goto L8
L17:
	;
	base.MemoryFill(m, v42, int32(0), v43)
	goto L8
L18:
	;
	v137 = int32(_a_F_fill_seq_fork_with_data_6)
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_fill_seq_fork_with_data[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_fill_seq_fork_with_data[2])) = v139 + int32(1)
	F_MarkBufferDirty(m, v23)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L2
	} else {
		goto L26
	}
L19:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_fill_seq_fork_with_data[3]))
	if v130 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v133 != 0 {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v135 = F_GetTopTransactionId(m)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L2
	} else {
		goto L25
	}
L23:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v134 != 0 {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	goto L18
L26:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v147 = int32(0)
	v149 = F_PageAddItemExtended(m, v42, v145, v146, v147, v147)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	if v149 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+118)))
	if v154 != int32(112) {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	goto L30
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L2
	} else {
		goto L46
	}
L31:
	;
	v195 = int32(_a_F_fill_seq_fork_with_data_6)
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_fill_seq_fork_with_data[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_fill_seq_fork_with_data[2])) = v197 - int32(1)
	F_UnlockReleaseBuffer(m, v23)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L2
	} else {
		goto L45
	}
L32:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L2
	} else {
		goto L40
	}
L33:
	;
	if l2 != int32(3) {
		goto L31
	} else {
		goto L39
	}
L34:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_fill_seq_fork_with_data[3]))
	if int32(0) < v158 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v161 != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	if l2 == int32(3) {
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v164 == int32(0) {
		goto L32
	} else {
		goto L38
	}
L38:
	;
	goto L31
L39:
	;
	goto L32
L40:
	;
	F_XLogRegisterBuffer(m, int32(0), v23, int32(6))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v175
	v177 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v177
	F_XLogRegisterData(m, v10+int32(24), int32(12))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_XLogRegisterData(m, v184, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v190 = F_XLogInsert(m, int32(15), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = base.I64_rotr(v190, int64(32))
	goto L31
L45:
	;
	m.G0 = v10 + int32(48)
	return
L46:
	;
	F_errmsg_internal(m, int32(_a_F_fill_seq_fork_with_data_7), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_fill_seq_fork_with_data_8), int32(405), int32(_a_F_fill_seq_fork_with_data_9))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
