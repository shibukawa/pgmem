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
	var v14 int32
	_ = v14
	var v17 int64
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
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int64
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int64
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = int64(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = l0
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v10)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v17
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
	if v42&int32(3) != 0 {
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
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v42+v84))) = int32(_a_F_fill_seq_fork_with_data_4)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = int32(2)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+20)))
	v94 = v92 | int32(768)
	*(*uint16)(unsafe.Add(mBase, uint32(v91)+20)) = uint16(v94)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v97 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v97
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96)+20)))
	v101 = v99 & int32(_a_F_fill_seq_fork_with_data_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v96)+20)) = uint16(v101)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v97
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+20)))
	v109 = v107 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v106)+20)) = uint16(v109)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v112 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v111)+16)) = uint16(v112)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = v97
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+118)))
	if v117 != int32(112) {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+10)) = int32(_a_F_fill_seq_fork_with_data_1)
	v75 = int32(_a_F_fill_seq_fork_with_data_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v42)+18)) = uint16(v75)
	v81 = int32(_a_F_fill_seq_fork_with_data_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v42)+16)) = uint16(v81)
	*(*uint16)(unsafe.Add(mBase, uint32(v42)+14)) = uint16(v81)
	goto L7
L9:
	;
	v69 = F___memset(m, v42, int32(0), int32(_a_F_fill_seq_fork_with_data_0))
	mBase = m.M
	goto L8
L10:
	;
	goto L9
L17:
	;
	v128 = int32(_a_F_fill_seq_fork_with_data_6)
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_fill_seq_fork_with_data[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_fill_seq_fork_with_data[2])) = v130 + int32(1)
	F_MarkBufferDirty(m, v23)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L2
	} else {
		goto L25
	}
L18:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_fill_seq_fork_with_data[3]))
	if v121 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v124 != 0 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v126 = F_GetTopTransactionId(m)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L24
	}
L22:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v125 != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L17
L25:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v138 = int32(0)
	v140 = F_PageAddItemExtended(m, v42, v136, v137, v138, v138)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	if v140 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+118)))
	if v145 != int32(112) {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	goto L29
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L2
	} else {
		goto L45
	}
L30:
	;
	v186 = int32(_a_F_fill_seq_fork_with_data_6)
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_fill_seq_fork_with_data[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_fill_seq_fork_with_data[2])) = v188 - int32(1)
	F_UnlockReleaseBuffer(m, v23)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L2
	} else {
		goto L44
	}
L31:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L39
	}
L32:
	;
	if l2 != int32(3) {
		goto L30
	} else {
		goto L38
	}
L33:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_fill_seq_fork_with_data[3]))
	if int32(0) < v149 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v152 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	if l2 == int32(3) {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v155 == int32(0) {
		goto L31
	} else {
		goto L37
	}
L37:
	;
	goto L30
L38:
	;
	goto L31
L39:
	;
	F_XLogRegisterBuffer(m, int32(0), v23, int32(6))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v166
	v168 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v168
	F_XLogRegisterData(m, v10+int32(24), int32(12))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_XLogRegisterData(m, v175, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v181 = F_XLogInsert(m, int32(15), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = base.I64_rotr(v181, int64(32))
	goto L30
L44:
	;
	m.G0 = v10 + int32(48)
	return
L45:
	;
	F_errmsg_internal(m, int32(_a_F_fill_seq_fork_with_data_7), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_fill_seq_fork_with_data_8), int32(405), int32(_a_F_fill_seq_fork_with_data_9))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
