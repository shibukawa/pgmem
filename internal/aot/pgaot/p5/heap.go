package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HeapTupleHeaderAdjustCmax(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	v4 = int32(0)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	if v5&int32(1) != 0 {
		v147 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v147)
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v8) < base.Ui32(int32(3)) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v128 == int32(0) {
		v147 = v4
		goto L1
	} else {
		goto L43
	}
L4:
	;
	v128 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderAdjustCmax[0]))
	if v19 == v8 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v128 = int32(1)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderAdjustCmax[1]))
	if v23 <= int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v128 = v120
	goto L3
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderAdjustCmax[2]))
	if v27 == int32(0) {
		v120 = int32(0)
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderAdjustCmax[3]))
	v91 = int32(0)
	v93 = v23 - int32(1)
	goto L33
L14:
	;
	v32 = v27
	goto L15
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	if v37 == int32(4) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v120 = int32(0)
	goto L10
L17:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v32)+80))
	if v84 != 0 {
		v32 = v84
		goto L15
	} else {
		goto L32
	}
L18:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v40 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v43 = int32(1)
	if v8 == v40 {
		v120 = v43
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	v47 = v45 - int32(1)
	if v47 < int32(0) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v52 = int32(0)
	v54 = v47
	goto L22
L22:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
	v60 = int32(2)
	v61 = base.I32_div_s(v54-v52, v60)
	v62 = v61 + v52
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v58+v62<<(uint(v60)%32))))
	if v66 == v8 {
		v120 = v43
		goto L10
	} else {
		goto L24
	}
L23:
	;
	goto L17
L24:
	;
	v70 = F_TransactionIdPrecedes(m, v66, v8)
	mBase = m.M
	if v70 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v71 = v62 + int32(1)
	goto L27
L26:
	;
	v71 = v52
	goto L27
L27:
	;
	if v70 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v74 = v54
	goto L30
L29:
	;
	v74 = v62 - int32(1)
	goto L30
L30:
	;
	if v71 <= v74 {
		v52 = v71
		v54 = v74
		goto L22
	} else {
		goto L31
	}
L31:
	;
	goto L23
L32:
	;
	goto L16
L33:
	;
	v98 = int32(2)
	v99 = base.I32_div_s(v93-v91, v98)
	v100 = v99 + v91
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v89+v100<<(uint(v98)%32))))
	v105 = base.B2i32(v104 == v8)
	if v104 == v8 {
		v120 = v105
		goto L10
	} else {
		goto L35
	}
L34:
	;
	v120 = v105
	goto L10
L35:
	;
	v108 = base.B2i32(base.Ui32(v104) < base.Ui32(v8))
	if base.Ui32(v104) < base.Ui32(v8) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v109 = v100 + int32(1)
	goto L38
L37:
	;
	v109 = v91
	goto L38
L38:
	;
	if base.Ui32(v104) < base.Ui32(v8) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v112 = v93
	goto L41
L40:
	;
	v112 = v100 - int32(1)
	goto L41
L41:
	;
	if v109 <= v112 {
		v91 = v109
		v93 = v112
		goto L33
	} else {
		goto L42
	}
L42:
	;
	goto L34
L43:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v132&int32(32) != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderAdjustCmax[4]))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v131<<(uint(int32(3))%32))))
	v141 = v140
	goto L46
L45:
	;
	v141 = v131
	goto L46
L46:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v143 = F_GetComboCommandId(m, v141, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	return
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v143
	v147 = int32(1)
	goto L1
}
func F_heap_abort_speculative(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int64
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
		v22 = F_ReadBuffer(m, l0, v17|v18<<(uint(int32(16))%32))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			if v22 < int32(0) {
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[0]))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v27+(v22^int32(-1))<<(uint(int32(2))%32))))
				v41 = v33
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[1]))
				v41 = v35 + v22<<(uint(int32(13))%32) + int32(-8192)
			}
			F_LockBuffer(m, v22, int32(2))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v46
				v49 = v41 + int32(20)
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v45<<(uint(int32(2))%32))))
				*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(base.Ui32(v53) >> (uint(int32(17)) % 32))
				v59 = v41 + v53&int32(_a_F_heap_abort_speculative_0)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v59
				v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				*(*uint16)(unsafe.Add(mBase, uint32(v13)+20)) = uint16(v61)
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v63
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
				if v15 == v65 {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+68))
					if v68 != int32(99) {
						v71 = F_isTempToastNamespace(m, v68)
						mBase = m.M
						v73 = v71
					} else {
						v73 = int32(1)
					}
					if v73 == int32(0) {
						v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+16)))
						if v76 != int32(_a_F_heap_abort_speculative_1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v240 = m.ExcPending
							if v240 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_heap_abort_speculative_2), int32(0))
								mBase = m.M
								v244 = m.ExcPending
								if v244 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_heap_abort_speculative_3), int32(_a_F_heap_abort_speculative_4), int32(_a_F_heap_abort_speculative_5))
									mBase = m.M
									v249 = m.ExcPending
									if v249 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v80 = v13 + int32(16)
							v81 = int32(_a_F_heap_abort_speculative_6)
							v83 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
							*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v83 + int32(1)
							v88 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[3]))
							v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+136))
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v90))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v88)) == int32(0) {
								v102 = base.B2i32(base.Ui32(v88) < base.Ui32(v90))
							} else {
								v102 = int32(base.Ui32(v88-v90) >> (uint(int32(31)) % 32))
							}
							v104 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[3]))
							if v102 != 0 {
								v105 = v90
							} else {
								v105 = v104
							}
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
							if v106 != 0 {
								if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v106))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v105)) == int32(0) {
									v118 = base.B2i32(base.Ui32(v105) < base.Ui32(v106))
								} else {
									v118 = int32(base.Ui32(v105-v106) >> (uint(int32(31)) % 32))
								}
								if v118 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v49))) = v105
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v49))) = v105
							}
							*(*int32)(unsafe.Add(mBase, uint32(v59))) = int32(0)
							v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+20)))
							v126 = v124 & int32(_a_F_heap_abort_speculative_7)
							*(*uint16)(unsafe.Add(mBase, uint32(v59)+20)) = uint16(v126)
							v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+18)))
							v130 = v128 & int32(_a_F_heap_abort_speculative_8)
							*(*uint16)(unsafe.Add(mBase, uint32(v59)+18)) = uint16(v130)
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
							*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = v132
							v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+4)))
							*(*uint16)(unsafe.Add(mBase, uint32(v59)+16)) = uint16(v134)
							F_MarkBufferDirty(m, v22)
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return
							} else {
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+118)))
								if v139 != int32(112) {
									v200 = int32(_a_F_heap_abort_speculative_6)
									v202 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
									*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v202 - int32(1)
									F_LockBuffer(m, v22, int32(0))
									mBase = m.M
									v208 = m.ExcPending
									if v208 != 0 {
										return
									} else {
										v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+20)))
										if v209&int32(4) != 0 {
											F_heap_toast_delete(m, l0, v13+int32(12), int32(1))
											mBase = m.M
											v216 = m.ExcPending
											if v216 != 0 {
												return
											} else {
												F_ReleaseBuffer(m, v22)
												mBase = m.M
												v218 = m.ExcPending
												if v218 != 0 {
													return
												} else {
													F_pgstat_count_heap_delete(m, l0)
													mBase = m.M
													v220 = m.ExcPending
													if v220 != 0 {
														return
													} else {
														m.G0 = v13 + int32(32)
														return
													}
												}
											}
										} else {
											F_ReleaseBuffer(m, v22)
											mBase = m.M
											v218 = m.ExcPending
											if v218 != 0 {
												return
											} else {
												F_pgstat_count_heap_delete(m, l0)
												mBase = m.M
												v220 = m.ExcPending
												if v220 != 0 {
													return
												} else {
													m.G0 = v13 + int32(32)
													return
												}
											}
										}
									}
								} else {
									v143 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[4]))
									if v143 <= int32(0) {
										v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										if v146 != 0 {
											v200 = int32(_a_F_heap_abort_speculative_6)
											v202 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
											*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v202 - int32(1)
											F_LockBuffer(m, v22, int32(0))
											mBase = m.M
											v208 = m.ExcPending
											if v208 != 0 {
												return
											} else {
												v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+20)))
												if v209&int32(4) != 0 {
													F_heap_toast_delete(m, l0, v13+int32(12), int32(1))
													mBase = m.M
													v216 = m.ExcPending
													if v216 != 0 {
														return
													} else {
														F_ReleaseBuffer(m, v22)
														mBase = m.M
														v218 = m.ExcPending
														if v218 != 0 {
															return
														} else {
															F_pgstat_count_heap_delete(m, l0)
															mBase = m.M
															v220 = m.ExcPending
															if v220 != 0 {
																return
															} else {
																m.G0 = v13 + int32(32)
																return
															}
														}
													}
												} else {
													F_ReleaseBuffer(m, v22)
													mBase = m.M
													v218 = m.ExcPending
													if v218 != 0 {
														return
													} else {
														F_pgstat_count_heap_delete(m, l0)
														mBase = m.M
														v220 = m.ExcPending
														if v220 != 0 {
															return
														} else {
															m.G0 = v13 + int32(32)
															return
														}
													}
												}
											}
										} else {
											v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											if v147 != 0 {
												v200 = int32(_a_F_heap_abort_speculative_6)
												v202 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
												*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v202 - int32(1)
												F_LockBuffer(m, v22, int32(0))
												mBase = m.M
												v208 = m.ExcPending
												if v208 != 0 {
													return
												} else {
													v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+20)))
													if v209&int32(4) != 0 {
														F_heap_toast_delete(m, l0, v13+int32(12), int32(1))
														mBase = m.M
														v216 = m.ExcPending
														if v216 != 0 {
															return
														} else {
															F_ReleaseBuffer(m, v22)
															mBase = m.M
															v218 = m.ExcPending
															if v218 != 0 {
																return
															} else {
																F_pgstat_count_heap_delete(m, l0)
																mBase = m.M
																v220 = m.ExcPending
																if v220 != 0 {
																	return
																} else {
																	m.G0 = v13 + int32(32)
																	return
																}
															}
														}
													} else {
														F_ReleaseBuffer(m, v22)
														mBase = m.M
														v218 = m.ExcPending
														if v218 != 0 {
															return
														} else {
															F_pgstat_count_heap_delete(m, l0)
															mBase = m.M
															v220 = m.ExcPending
															if v220 != 0 {
																return
															} else {
																m.G0 = v13 + int32(32)
																return
															}
														}
													}
												}
											} else {
												v148 = int32(8)
												*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)) = uint8(v148)
												v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+18)))
												v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+20)))
												v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+20)))
												*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)) = uint16(v152)
												*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v15
												v159 = int32(1)
												v163 = int32(4)
												v178 = int32(base.Ui32(v150)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v151)>>(uint(v159)%32))&v148 | (int32(base.Ui32(v151)>>(uint(v163)%32))&v163 | (int32(base.Ui32(v151)>>(uint(int32(12))%32))&v159 | int32(base.Ui32(v151)>>(uint(int32(6))%32))&int32(2))))
												*(*uint8)(unsafe.Add(mBase, uint32(v13)+10)) = uint8(v178)
												F_XLogBeginInsert(m)
												mBase = m.M
												v181 = m.ExcPending
												if v181 != 0 {
													return
												} else {
													F_XLogRegisterData(m, v13+int32(4), int32(8))
													mBase = m.M
													v186 = m.ExcPending
													if v186 != 0 {
														return
													} else {
														F_XLogRegisterBuffer(m, int32(0), v22, int32(8))
														mBase = m.M
														v190 = m.ExcPending
														if v190 != 0 {
															return
														} else {
															v193 = F_XLogInsert(m, int32(10), int32(16))
															mBase = m.M
															v194 = m.ExcPending
															if v194 != 0 {
																return
															} else {
																*(*int64)(unsafe.Add(mBase, uint32(v41))) = base.I64_rotr(v193, int64(32))
																v200 = int32(_a_F_heap_abort_speculative_6)
																v202 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
																*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v202 - int32(1)
																F_LockBuffer(m, v22, int32(0))
																mBase = m.M
																v208 = m.ExcPending
																if v208 != 0 {
																	return
																} else {
																	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+20)))
																	if v209&int32(4) != 0 {
																		F_heap_toast_delete(m, l0, v13+int32(12), int32(1))
																		mBase = m.M
																		v216 = m.ExcPending
																		if v216 != 0 {
																			return
																		} else {
																			F_ReleaseBuffer(m, v22)
																			mBase = m.M
																			v218 = m.ExcPending
																			if v218 != 0 {
																				return
																			} else {
																				F_pgstat_count_heap_delete(m, l0)
																				mBase = m.M
																				v220 = m.ExcPending
																				if v220 != 0 {
																					return
																				} else {
																					m.G0 = v13 + int32(32)
																					return
																				}
																			}
																		}
																	} else {
																		F_ReleaseBuffer(m, v22)
																		mBase = m.M
																		v218 = m.ExcPending
																		if v218 != 0 {
																			return
																		} else {
																			F_pgstat_count_heap_delete(m, l0)
																			mBase = m.M
																			v220 = m.ExcPending
																			if v220 != 0 {
																				return
																			} else {
																				m.G0 = v13 + int32(32)
																				return
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										v148 = int32(8)
										*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)) = uint8(v148)
										v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+18)))
										v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+20)))
										v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+20)))
										*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)) = uint16(v152)
										*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v15
										v159 = int32(1)
										v163 = int32(4)
										v178 = int32(base.Ui32(v150)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v151)>>(uint(v159)%32))&v148 | (int32(base.Ui32(v151)>>(uint(v163)%32))&v163 | (int32(base.Ui32(v151)>>(uint(int32(12))%32))&v159 | int32(base.Ui32(v151)>>(uint(int32(6))%32))&int32(2))))
										*(*uint8)(unsafe.Add(mBase, uint32(v13)+10)) = uint8(v178)
										F_XLogBeginInsert(m)
										mBase = m.M
										v181 = m.ExcPending
										if v181 != 0 {
											return
										} else {
											F_XLogRegisterData(m, v13+int32(4), int32(8))
											mBase = m.M
											v186 = m.ExcPending
											if v186 != 0 {
												return
											} else {
												F_XLogRegisterBuffer(m, int32(0), v22, int32(8))
												mBase = m.M
												v190 = m.ExcPending
												if v190 != 0 {
													return
												} else {
													v193 = F_XLogInsert(m, int32(10), int32(16))
													mBase = m.M
													v194 = m.ExcPending
													if v194 != 0 {
														return
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v41))) = base.I64_rotr(v193, int64(32))
														v200 = int32(_a_F_heap_abort_speculative_6)
														v202 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
														*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v202 - int32(1)
														F_LockBuffer(m, v22, int32(0))
														mBase = m.M
														v208 = m.ExcPending
														if v208 != 0 {
															return
														} else {
															v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+20)))
															if v209&int32(4) != 0 {
																F_heap_toast_delete(m, l0, v13+int32(12), int32(1))
																mBase = m.M
																v216 = m.ExcPending
																if v216 != 0 {
																	return
																} else {
																	F_ReleaseBuffer(m, v22)
																	mBase = m.M
																	v218 = m.ExcPending
																	if v218 != 0 {
																		return
																	} else {
																		F_pgstat_count_heap_delete(m, l0)
																		mBase = m.M
																		v220 = m.ExcPending
																		if v220 != 0 {
																			return
																		} else {
																			m.G0 = v13 + int32(32)
																			return
																		}
																	}
																}
															} else {
																F_ReleaseBuffer(m, v22)
																mBase = m.M
																v218 = m.ExcPending
																if v218 != 0 {
																	return
																} else {
																	F_pgstat_count_heap_delete(m, l0)
																	mBase = m.M
																	v220 = m.ExcPending
																	if v220 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(32)
																		return
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v80 = v13 + int32(16)
						v81 = int32(_a_F_heap_abort_speculative_6)
						v83 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
						*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v83 + int32(1)
						v88 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[3]))
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+136))
						if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v90))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v88)) == int32(0) {
							v102 = base.B2i32(base.Ui32(v88) < base.Ui32(v90))
						} else {
							v102 = int32(base.Ui32(v88-v90) >> (uint(int32(31)) % 32))
						}
						v104 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[3]))
						if v102 != 0 {
							v105 = v90
						} else {
							v105 = v104
						}
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
						if v106 != 0 {
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v106))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v105)) == int32(0) {
								v118 = base.B2i32(base.Ui32(v105) < base.Ui32(v106))
							} else {
								v118 = int32(base.Ui32(v105-v106) >> (uint(int32(31)) % 32))
							}
							if v118 == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v49))) = v105
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v49))) = v105
						}
						*(*int32)(unsafe.Add(mBase, uint32(v59))) = int32(0)
						v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+20)))
						v126 = v124 & int32(_a_F_heap_abort_speculative_7)
						*(*uint16)(unsafe.Add(mBase, uint32(v59)+20)) = uint16(v126)
						v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+18)))
						v130 = v128 & int32(_a_F_heap_abort_speculative_8)
						*(*uint16)(unsafe.Add(mBase, uint32(v59)+18)) = uint16(v130)
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
						*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = v132
						v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+4)))
						*(*uint16)(unsafe.Add(mBase, uint32(v59)+16)) = uint16(v134)
						F_MarkBufferDirty(m, v22)
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return
						} else {
							v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+118)))
							if v139 != int32(112) {
								v200 = int32(_a_F_heap_abort_speculative_6)
								v202 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
								*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v202 - int32(1)
								F_LockBuffer(m, v22, int32(0))
								mBase = m.M
								v208 = m.ExcPending
								if v208 != 0 {
									return
								} else {
									v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+20)))
									if v209&int32(4) != 0 {
										F_heap_toast_delete(m, l0, v13+int32(12), int32(1))
										mBase = m.M
										v216 = m.ExcPending
										if v216 != 0 {
											return
										} else {
											F_ReleaseBuffer(m, v22)
											mBase = m.M
											v218 = m.ExcPending
											if v218 != 0 {
												return
											} else {
												F_pgstat_count_heap_delete(m, l0)
												mBase = m.M
												v220 = m.ExcPending
												if v220 != 0 {
													return
												} else {
													m.G0 = v13 + int32(32)
													return
												}
											}
										}
									} else {
										F_ReleaseBuffer(m, v22)
										mBase = m.M
										v218 = m.ExcPending
										if v218 != 0 {
											return
										} else {
											F_pgstat_count_heap_delete(m, l0)
											mBase = m.M
											v220 = m.ExcPending
											if v220 != 0 {
												return
											} else {
												m.G0 = v13 + int32(32)
												return
											}
										}
									}
								}
							} else {
								v143 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[4]))
								if v143 <= int32(0) {
									v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									if v146 != 0 {
										v200 = int32(_a_F_heap_abort_speculative_6)
										v202 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
										*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v202 - int32(1)
										F_LockBuffer(m, v22, int32(0))
										mBase = m.M
										v208 = m.ExcPending
										if v208 != 0 {
											return
										} else {
											v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+20)))
											if v209&int32(4) != 0 {
												F_heap_toast_delete(m, l0, v13+int32(12), int32(1))
												mBase = m.M
												v216 = m.ExcPending
												if v216 != 0 {
													return
												} else {
													F_ReleaseBuffer(m, v22)
													mBase = m.M
													v218 = m.ExcPending
													if v218 != 0 {
														return
													} else {
														F_pgstat_count_heap_delete(m, l0)
														mBase = m.M
														v220 = m.ExcPending
														if v220 != 0 {
															return
														} else {
															m.G0 = v13 + int32(32)
															return
														}
													}
												}
											} else {
												F_ReleaseBuffer(m, v22)
												mBase = m.M
												v218 = m.ExcPending
												if v218 != 0 {
													return
												} else {
													F_pgstat_count_heap_delete(m, l0)
													mBase = m.M
													v220 = m.ExcPending
													if v220 != 0 {
														return
													} else {
														m.G0 = v13 + int32(32)
														return
													}
												}
											}
										}
									} else {
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										if v147 != 0 {
											v200 = int32(_a_F_heap_abort_speculative_6)
											v202 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
											*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v202 - int32(1)
											F_LockBuffer(m, v22, int32(0))
											mBase = m.M
											v208 = m.ExcPending
											if v208 != 0 {
												return
											} else {
												v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+20)))
												if v209&int32(4) != 0 {
													F_heap_toast_delete(m, l0, v13+int32(12), int32(1))
													mBase = m.M
													v216 = m.ExcPending
													if v216 != 0 {
														return
													} else {
														F_ReleaseBuffer(m, v22)
														mBase = m.M
														v218 = m.ExcPending
														if v218 != 0 {
															return
														} else {
															F_pgstat_count_heap_delete(m, l0)
															mBase = m.M
															v220 = m.ExcPending
															if v220 != 0 {
																return
															} else {
																m.G0 = v13 + int32(32)
																return
															}
														}
													}
												} else {
													F_ReleaseBuffer(m, v22)
													mBase = m.M
													v218 = m.ExcPending
													if v218 != 0 {
														return
													} else {
														F_pgstat_count_heap_delete(m, l0)
														mBase = m.M
														v220 = m.ExcPending
														if v220 != 0 {
															return
														} else {
															m.G0 = v13 + int32(32)
															return
														}
													}
												}
											}
										} else {
											v148 = int32(8)
											*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)) = uint8(v148)
											v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+18)))
											v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+20)))
											v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+20)))
											*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)) = uint16(v152)
											*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v15
											v159 = int32(1)
											v163 = int32(4)
											v178 = int32(base.Ui32(v150)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v151)>>(uint(v159)%32))&v148 | (int32(base.Ui32(v151)>>(uint(v163)%32))&v163 | (int32(base.Ui32(v151)>>(uint(int32(12))%32))&v159 | int32(base.Ui32(v151)>>(uint(int32(6))%32))&int32(2))))
											*(*uint8)(unsafe.Add(mBase, uint32(v13)+10)) = uint8(v178)
											F_XLogBeginInsert(m)
											mBase = m.M
											v181 = m.ExcPending
											if v181 != 0 {
												return
											} else {
												F_XLogRegisterData(m, v13+int32(4), int32(8))
												mBase = m.M
												v186 = m.ExcPending
												if v186 != 0 {
													return
												} else {
													F_XLogRegisterBuffer(m, int32(0), v22, int32(8))
													mBase = m.M
													v190 = m.ExcPending
													if v190 != 0 {
														return
													} else {
														v193 = F_XLogInsert(m, int32(10), int32(16))
														mBase = m.M
														v194 = m.ExcPending
														if v194 != 0 {
															return
														} else {
															*(*int64)(unsafe.Add(mBase, uint32(v41))) = base.I64_rotr(v193, int64(32))
															v200 = int32(_a_F_heap_abort_speculative_6)
															v202 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
															*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v202 - int32(1)
															F_LockBuffer(m, v22, int32(0))
															mBase = m.M
															v208 = m.ExcPending
															if v208 != 0 {
																return
															} else {
																v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+20)))
																if v209&int32(4) != 0 {
																	F_heap_toast_delete(m, l0, v13+int32(12), int32(1))
																	mBase = m.M
																	v216 = m.ExcPending
																	if v216 != 0 {
																		return
																	} else {
																		F_ReleaseBuffer(m, v22)
																		mBase = m.M
																		v218 = m.ExcPending
																		if v218 != 0 {
																			return
																		} else {
																			F_pgstat_count_heap_delete(m, l0)
																			mBase = m.M
																			v220 = m.ExcPending
																			if v220 != 0 {
																				return
																			} else {
																				m.G0 = v13 + int32(32)
																				return
																			}
																		}
																	}
																} else {
																	F_ReleaseBuffer(m, v22)
																	mBase = m.M
																	v218 = m.ExcPending
																	if v218 != 0 {
																		return
																	} else {
																		F_pgstat_count_heap_delete(m, l0)
																		mBase = m.M
																		v220 = m.ExcPending
																		if v220 != 0 {
																			return
																		} else {
																			m.G0 = v13 + int32(32)
																			return
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								} else {
									v148 = int32(8)
									*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)) = uint8(v148)
									v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+18)))
									v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+20)))
									v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+20)))
									*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)) = uint16(v152)
									*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v15
									v159 = int32(1)
									v163 = int32(4)
									v178 = int32(base.Ui32(v150)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v151)>>(uint(v159)%32))&v148 | (int32(base.Ui32(v151)>>(uint(v163)%32))&v163 | (int32(base.Ui32(v151)>>(uint(int32(12))%32))&v159 | int32(base.Ui32(v151)>>(uint(int32(6))%32))&int32(2))))
									*(*uint8)(unsafe.Add(mBase, uint32(v13)+10)) = uint8(v178)
									F_XLogBeginInsert(m)
									mBase = m.M
									v181 = m.ExcPending
									if v181 != 0 {
										return
									} else {
										F_XLogRegisterData(m, v13+int32(4), int32(8))
										mBase = m.M
										v186 = m.ExcPending
										if v186 != 0 {
											return
										} else {
											F_XLogRegisterBuffer(m, int32(0), v22, int32(8))
											mBase = m.M
											v190 = m.ExcPending
											if v190 != 0 {
												return
											} else {
												v193 = F_XLogInsert(m, int32(10), int32(16))
												mBase = m.M
												v194 = m.ExcPending
												if v194 != 0 {
													return
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v41))) = base.I64_rotr(v193, int64(32))
													v200 = int32(_a_F_heap_abort_speculative_6)
													v202 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
													*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v202 - int32(1)
													F_LockBuffer(m, v22, int32(0))
													mBase = m.M
													v208 = m.ExcPending
													if v208 != 0 {
														return
													} else {
														v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+20)))
														if v209&int32(4) != 0 {
															F_heap_toast_delete(m, l0, v13+int32(12), int32(1))
															mBase = m.M
															v216 = m.ExcPending
															if v216 != 0 {
																return
															} else {
																F_ReleaseBuffer(m, v22)
																mBase = m.M
																v218 = m.ExcPending
																if v218 != 0 {
																	return
																} else {
																	F_pgstat_count_heap_delete(m, l0)
																	mBase = m.M
																	v220 = m.ExcPending
																	if v220 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(32)
																		return
																	}
																}
															}
														} else {
															F_ReleaseBuffer(m, v22)
															mBase = m.M
															v218 = m.ExcPending
															if v218 != 0 {
																return
															} else {
																F_pgstat_count_heap_delete(m, l0)
																mBase = m.M
																v220 = m.ExcPending
																if v220 != 0 {
																	return
																} else {
																	m.G0 = v13 + int32(32)
																	return
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v227 = m.ExcPending
					if v227 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_heap_abort_speculative_9), int32(0))
						mBase = m.M
						v231 = m.ExcPending
						if v231 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_heap_abort_speculative_3), int32(_a_F_heap_abort_speculative_10), int32(_a_F_heap_abort_speculative_5))
							mBase = m.M
							v236 = m.ExcPending
							if v236 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_heap_beginscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	F_RelationIncrementReferenceCount(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = F_palloc(m, int32(692))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = l5
			*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
			*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = int64(0)
			if l1 == int32(0) {
				v26 = l5 & int32(-257)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v26
				v28 = v26
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				switch v24 {
				case 0, 5:
					v28 = l5
				default:
					v26 = l5 & int32(-257)
					*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v26
					v28 = v26
				}
			}
			if v28&int32(5) != 0 {
				F_PredicateLockRelation(m, l0, l1)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v33
					if l4 != 0 {
						v36 = F_palloc(m, int32(16))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v39 = v36
							*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v39
							if int32(0) < l2 {
								v45 = F_palloc(m, l2*int32(48))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									v48 = v45
									*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v48
									F_initscan(m, v13, l3, int32(0))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = int32(0)
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
										if v55&int32(17) != 0 {
											v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
											v60 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
											if v64 != 0 {
												v65 = int32(138)
											} else {
												v65 = int32(139)
											}
											v67 = F_read_stream_begin_relation(m, int32(10), v59, v60, int32(0), v65, v13, int32(0))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												v81 = v67
												*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v81
												return v13
											}
										} else {
											if v55&int32(2) == int32(0) {
												return v13
											} else {
												v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
												v79 = F_read_stream_begin_relation(m, int32(8), v74, v75, int32(0), int32(140), v13, int32(12))
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													v81 = v79
													*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v81
													return v13
												}
											}
										}
									}
								}
							} else {
								v48 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v48
								F_initscan(m, v13, l3, int32(0))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = int32(0)
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
									if v55&int32(17) != 0 {
										v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
										if v64 != 0 {
											v65 = int32(138)
										} else {
											v65 = int32(139)
										}
										v67 = F_read_stream_begin_relation(m, int32(10), v59, v60, int32(0), v65, v13, int32(0))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return int32(0)
										} else {
											v81 = v67
											*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v81
											return v13
										}
									} else {
										if v55&int32(2) == int32(0) {
											return v13
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
											v79 = F_read_stream_begin_relation(m, int32(8), v74, v75, int32(0), int32(140), v13, int32(12))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int32(0)
											} else {
												v81 = v79
												*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v81
												return v13
											}
										}
									}
								}
							}
						}
					} else {
						v39 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v39
						if int32(0) < l2 {
							v45 = F_palloc(m, l2*int32(48))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v48 = v45
								*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v48
								F_initscan(m, v13, l3, int32(0))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = int32(0)
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
									if v55&int32(17) != 0 {
										v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
										if v64 != 0 {
											v65 = int32(138)
										} else {
											v65 = int32(139)
										}
										v67 = F_read_stream_begin_relation(m, int32(10), v59, v60, int32(0), v65, v13, int32(0))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return int32(0)
										} else {
											v81 = v67
											*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v81
											return v13
										}
									} else {
										if v55&int32(2) == int32(0) {
											return v13
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
											v79 = F_read_stream_begin_relation(m, int32(8), v74, v75, int32(0), int32(140), v13, int32(12))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int32(0)
											} else {
												v81 = v79
												*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v81
												return v13
											}
										}
									}
								}
							}
						} else {
							v48 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v48
							F_initscan(m, v13, l3, int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = int32(0)
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
								if v55&int32(17) != 0 {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
									if v64 != 0 {
										v65 = int32(138)
									} else {
										v65 = int32(139)
									}
									v67 = F_read_stream_begin_relation(m, int32(10), v59, v60, int32(0), v65, v13, int32(0))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v81 = v67
										*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v81
										return v13
									}
								} else {
									if v55&int32(2) == int32(0) {
										return v13
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										v79 = F_read_stream_begin_relation(m, int32(8), v74, v75, int32(0), int32(140), v13, int32(12))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											v81 = v79
											*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v81
											return v13
										}
									}
								}
							}
						}
					}
				}
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v33
				if l4 != 0 {
					v36 = F_palloc(m, int32(16))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v39 = v36
						*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v39
						if int32(0) < l2 {
							v45 = F_palloc(m, l2*int32(48))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v48 = v45
								*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v48
								F_initscan(m, v13, l3, int32(0))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = int32(0)
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
									if v55&int32(17) != 0 {
										v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
										if v64 != 0 {
											v65 = int32(138)
										} else {
											v65 = int32(139)
										}
										v67 = F_read_stream_begin_relation(m, int32(10), v59, v60, int32(0), v65, v13, int32(0))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return int32(0)
										} else {
											v81 = v67
											*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v81
											return v13
										}
									} else {
										if v55&int32(2) == int32(0) {
											return v13
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
											v79 = F_read_stream_begin_relation(m, int32(8), v74, v75, int32(0), int32(140), v13, int32(12))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int32(0)
											} else {
												v81 = v79
												*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v81
												return v13
											}
										}
									}
								}
							}
						} else {
							v48 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v48
							F_initscan(m, v13, l3, int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = int32(0)
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
								if v55&int32(17) != 0 {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
									if v64 != 0 {
										v65 = int32(138)
									} else {
										v65 = int32(139)
									}
									v67 = F_read_stream_begin_relation(m, int32(10), v59, v60, int32(0), v65, v13, int32(0))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v81 = v67
										*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v81
										return v13
									}
								} else {
									if v55&int32(2) == int32(0) {
										return v13
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										v79 = F_read_stream_begin_relation(m, int32(8), v74, v75, int32(0), int32(140), v13, int32(12))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											v81 = v79
											*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v81
											return v13
										}
									}
								}
							}
						}
					}
				} else {
					v39 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v39
					if int32(0) < l2 {
						v45 = F_palloc(m, l2*int32(48))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v48 = v45
							*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v48
							F_initscan(m, v13, l3, int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = int32(0)
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
								if v55&int32(17) != 0 {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
									if v64 != 0 {
										v65 = int32(138)
									} else {
										v65 = int32(139)
									}
									v67 = F_read_stream_begin_relation(m, int32(10), v59, v60, int32(0), v65, v13, int32(0))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v81 = v67
										*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v81
										return v13
									}
								} else {
									if v55&int32(2) == int32(0) {
										return v13
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										v79 = F_read_stream_begin_relation(m, int32(8), v74, v75, int32(0), int32(140), v13, int32(12))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											v81 = v79
											*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v81
											return v13
										}
									}
								}
							}
						}
					} else {
						v48 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v48
						F_initscan(m, v13, l3, int32(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = int32(0)
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
							if v55&int32(17) != 0 {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
								if v64 != 0 {
									v65 = int32(138)
								} else {
									v65 = int32(139)
								}
								v67 = F_read_stream_begin_relation(m, int32(10), v59, v60, int32(0), v65, v13, int32(0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v81 = v67
									*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v81
									return v13
								}
							} else {
								if v55&int32(2) == int32(0) {
									return v13
								} else {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
									v79 = F_read_stream_begin_relation(m, int32(8), v74, v75, int32(0), int32(140), v13, int32(12))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										v81 = v79
										*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v81
										return v13
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_heap_compare_slots_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	v4 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+112))
	if v11 <= v4 {
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+120))
	v17 = int32(2)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+l1<<(uint(v17)%32))))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16+l0<<(uint(v17)%32))))
	v30 = v4
	goto L6
L4:
	;
	return int32(-1)
L5:
	;
	v99 = int32(0)
	if v90 < v99 {
		goto L34
	} else {
		goto L35
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+116))
	v38 = v35 + v30*int32(36)
	v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+10)))
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+6)))
	if v40 < v39 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	F_slot_getsomeattrs_int(m, v24, v39)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v47 = v39 - int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+v48))))
	v52 = v47 << (uint(int32(2)) % 32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52+v53)))
	v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+6)))
	if v56 < v39 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	return int32(0)
L12:
	;
	goto L10
L13:
	;
	F_slot_getsomeattrs_int(m, v20, v39)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v47))))
	if v50&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v93 = v30 + int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+112))
	if v93 < v94 {
		v30 = v93
		goto L6
	} else {
		goto L33
	}
L18:
	;
	if v62&int32(1) != 0 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v62&int32(1) != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+9)))
	if v67 == int32(0) {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	return int32(1)
L23:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+9)))
	if v74 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77+v52)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v81 = m.T0[v80].(func(*base.Module, int32, int32, int32) int32)(m, v55, v79, v38)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L11
	} else {
		goto L27
	}
L26:
	;
	return int32(1)
L27:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+8)))
	if v83 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v81 < int32(0) {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	v90 = v81
	goto L30
L30:
	;
	if v90 != 0 {
		goto L5
	} else {
		goto L32
	}
L31:
	;
	v90 = int32(0) - v81
	goto L30
L32:
	;
	goto L17
L33:
	;
	goto L7
L34:
	;
	v103 = int32(1)
	goto L36
L35:
	;
	v103 = v99 - v90
	goto L36
L36:
	;
	return v103
}
func F_heap_compute_data_size(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	v4 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 <= v4 {
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
	v18 = int32(0)
	v23 = v4
	goto L4
L4:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+l2))))
	if v28 != 0 {
		v126 = v23
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return v126
L6:
	;
	v129 = v18 + int32(1)
	if v129 != v10 {
		v18 = v129
		v23 = v126
		goto L4
	} else {
		goto L31
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1+v18<<(uint(int32(2))%32))))
	v35 = l0 + int32(20) + v18<<(uint(int32(4))%32)
	v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35)+4)))
	if v36 == int32(-1) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v126 = v119 + v121
	goto L6
L9:
	;
	if v55&int32(1) != 0 {
		goto L28
	} else {
		goto L29
	}
L10:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	if base.Ui32((v91-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v119 = int32(6)
		v121 = v23
		goto L8
	} else {
		goto L24
	}
L11:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+7)))
	if v39 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+12)))
	v82 = int32(0)
	v84 = (v23 + v78 - int32(1)) & (v82 - v78)
	if v82 < v36 {
		v119 = v36
		v121 = v84
		goto L8
	} else {
		goto L23
	}
L14:
	;
	if v55&int32(255) != int32(1) {
		goto L9
	} else {
		goto L19
	}
L15:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	v55 = v54
	goto L14
L16:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v42&int32(3) != 0 {
		v55 = v42
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v49 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(3)
	if base.Ui32(int32(127)) < base.Ui32(v49) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v126 = v49 + v23
	goto L6
L19:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	if v60&int32(254) != int32(2) {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+12)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v32)+2))
	v73 = F_EOH_get_flat_size(m, v72)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	v126 = (v23+v65-int32(1))&(int32(0)-v65) + v73
	goto L6
L23:
	;
	v87 = F_strlen(m, v32)
	mBase = m.M
	v119 = v87 + int32(1)
	v121 = v84
	goto L8
L24:
	;
	v98 = int32(18)
	if v91 == v98 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v102 = v98
	goto L27
L26:
	;
	v102 = int32(2)
	goto L27
L27:
	;
	v119 = v102
	v121 = v23
	goto L8
L28:
	;
	v119 = int32(base.Ui32(v55&int32(254)) >> (uint(int32(1)) % 32))
	v121 = v23
	goto L8
L29:
	;
	goto L30
L30:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+12)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v119 = int32(base.Ui32(v116) >> (uint(int32(2)) % 32))
	v121 = (v23 + v109 - int32(1)) & (int32(0) - v109)
	goto L8
L31:
	;
	goto L5
}
func F_heap_copy_minimal_tuple(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = F_palloc(m, v4+l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if l1 != 0 {
			base.MemoryFill(m, v6, int32(0), l1)
		} else {
		}
		v12 = l1 + v6
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v13 != 0 {
			base.MemoryCopy(m, v12, l0, v13)
		} else {
		}
		return v12
	}
}
func F_heap_endscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ReleaseBuffer(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v6 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	F_read_stream_end(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_RelationDecrementReferenceCount(m, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v12 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	F_pfree(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v15 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	F_bms_free(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v18 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	F_pfree(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v21&int32(2) != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_UnregisterSnapshot(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	F_pfree(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	return
}
func F_heap_entry_is_visible(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = F_table_slot_create(m, v5, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_heap_entry_is_visible[0]))
		v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_entry_is_visible[1])))
		if v15&int32(1) != 0 {
			v18 = int32(0)
		} else {
			v18 = v13
		}
		if v18 == int32(0) {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+188))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
			v25 = m.T0[v24].(func(*base.Module, int32, int32, int32, int32) int32)(m, v21, l1, v22, v7)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v7 != 0 {
					F_ExecDropSingleTupleTableSlot(m, v7)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						return v25
					}
				} else {
					return v25
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_heap_entry_is_visible_0), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_heap_entry_is_visible_1), int32(1264), int32(_a_F_heap_entry_is_visible_2))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_heap_fetch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	v15 = F_ReadBuffer(m, l0, v10|v11<<(uint(int32(16))%32))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		F_LockBuffer(m, v15, int32(1))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v15 < int32(0) {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch[0]))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v25+(v15^int32(-1))<<(uint(int32(2))%32))))
				v39 = v31
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_heap_fetch[1]))
				v39 = v33 + v15<<(uint(int32(13))%32) + int32(-8192)
			}
			v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
			if v40 == int32(0) {
				F_LockBuffer(m, v15, int32(0))
				mBase = m.M
				v114 = m.ExcPending
				if v114 != 0 {
					return int32(0)
				} else {
					F_ReleaseBuffer(m, v15)
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						v119 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v119
						*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v119
						return v119
					}
				}
			} else {
				v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+12)))
				if base.Ui32(int32(25)) <= base.Ui32(v43) {
					v51 = int32(base.Ui32(v43+int32(_a_F_heap_fetch_0)) >> (uint(int32(2)) % 32))
				} else {
					v51 = int32(0)
				}
				if base.Ui32(v51&int32(_a_F_heap_fetch_1)) < base.Ui32(v40) {
					F_LockBuffer(m, v15, int32(0))
					mBase = m.M
					v114 = m.ExcPending
					if v114 != 0 {
						return int32(0)
					} else {
						F_ReleaseBuffer(m, v15)
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return int32(0)
						} else {
							v119 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v119
							*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v119
							return v119
						}
					}
				} else {
					v59 = v39 + v40<<(uint(int32(2))%32) + int32(20)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
					if v60&int32(_a_F_heap_fetch_2) != int32(_a_F_heap_fetch_3) {
						F_LockBuffer(m, v15, int32(0))
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return int32(0)
						} else {
							F_ReleaseBuffer(m, v15)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								v119 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v119
								*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v119
								return v119
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v39 + v60&int32(_a_F_heap_fetch_4)
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(base.Ui32(v69) >> (uint(int32(17)) % 32))
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v73
						v75 = F_HeapTupleSatisfiesVisibility(m, l2, l1, v15)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							if v75 != 0 {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
								v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+20)))
								v81 = int32(768)
								if v80&v81 != v81 {
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
									v87 = v85
								} else {
									v87 = int32(2)
								}
								F_PredicateLockTID(m, l0, l2+int32(4), l1, v87)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									F_HeapCheckForSerializableConflictOut(m, int32(1), l0, l2, v15, l1)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										F_LockBuffer(m, v15, int32(0))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v15
											return int32(1)
										}
									}
								}
							} else {
								F_HeapCheckForSerializableConflictOut(m, int32(0), l0, l2, v15, l1)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									F_LockBuffer(m, v15, int32(0))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int32(0)
									} else {
										if l4 == int32(0) {
											F_ReleaseBuffer(m, v15)
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return int32(0)
											} else {
												v119 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = v119
												*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v119
												return v119
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v15
											return int32(0)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_heap_get_latest_tid(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v19 = v16 + int32(12)
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = int32(0)
	v31 = v20
	v32 = v21
	v33 = v22
	goto L1
L1:
	;
	v41 = F_ReadBuffer(m, v24, v33<<(uint(int32(16))%32)|v32)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	F_UnlockReleaseBuffer(m, v41)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L3
	} else {
		goto L59
	}
L3:
	;
	return
L4:
	;
	F_LockBuffer(m, v41, int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v41 < int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v31 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_heap_get_latest_tid[0]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v41^int32(-1))<<(uint(int32(2))%32))))
	v63 = v55
	goto L6
L8:
	;
	goto L9
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_heap_get_latest_tid[1]))
	v63 = v57 + v41<<(uint(int32(13))%32) + int32(-8192)
	goto L6
L10:
	;
	goto L2
L11:
	;
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v66) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v74 = int32(base.Ui32(v66+int32(_a_F_heap_get_latest_tid_0)) >> (uint(int32(2)) % 32))
	goto L14
L13:
	;
	v74 = int32(0)
	goto L14
L14:
	;
	if base.Ui32(v74&int32(_a_F_heap_get_latest_tid_1)) < base.Ui32(v31) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v82 = v63 + v31<<(uint(int32(2))%32) + int32(20)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v83&int32(_a_F_heap_get_latest_tid_2) != int32(_a_F_heap_get_latest_tid_3) {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+16)) = uint16(v31)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+14)) = uint16(v32)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+12)) = uint16(v33)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v94 = v63 + v91&int32(_a_F_heap_get_latest_tid_4)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(base.Ui32(v96) >> (uint(int32(17)) % 32))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v100
	if v28 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+20)))
	v103 = int32(768)
	if v102&v103 != v103 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v112 = v16 + int32(8)
	v113 = F_HeapTupleSatisfiesVisibility(m, v112, v23, v41)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L3
	} else {
		goto L24
	}
L20:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v109 = v107
	goto L22
L21:
	;
	v109 = int32(2)
	goto L22
L22:
	;
	if v109 != v28 {
		goto L10
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	F_HeapCheckForSerializableConflictOut(m, v113, v24, v112, v41, v23)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	if v113 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v31)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)) = uint16(v32)
	*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v33)
	goto L28
L27:
	;
	goto L28
L28:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+21)))
	if v121&int32(8) != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	v124 = F_HeapTupleHeaderIsOnlyLocked(m, v120)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	if v124 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+16)))
	if v127 == int32(_a_F_heap_get_latest_tid_5) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+12)))
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+14)))
	if v130&v131 == int32(_a_F_heap_get_latest_tid_1) {
		goto L10
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v136 = v126 + int32(12)
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+2)))
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19))))
	v139 = int32(16)
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+2)))
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136))))
	if v137|v138<<(uint(v139)%32) == v142|v143<<(uint(v139)%32) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	goto L34
L36:
	;
	if v153 != 0 {
		goto L10
	} else {
		goto L42
	}
L37:
	;
	goto L36
L38:
	;
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+4)))
	if v149 == v150 {
		v153 = int32(1)
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v153 = int32(0)
	goto L37
L41:
	;
	goto L40
L42:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+16)))
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+14)))
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+12)))
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+20)))
	if v159&int32(_a_F_heap_get_latest_tid_6) != int32(_a_F_heap_get_latest_tid_7) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_UnlockReleaseBuffer(m, v41)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L3
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v166 = int32(0)
	v170 = F_GetMultiXactIdMembers(m, v155, v16+int32(28), v166)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L3
	} else {
		goto L47
	}
L46:
	;
	v28 = v155
	v31 = v156
	v32 = v157
	v33 = v158
	goto L1
L47:
	;
	if int32(0) < v170 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v176 = int32(0)
	goto L53
L49:
	;
	v207 = v166
	goto L50
L50:
	;
	F_UnlockReleaseBuffer(m, v41)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L3
	} else {
		goto L58
	}
L51:
	;
	F_pfree(m, v175)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L3
	} else {
		goto L57
	}
L52:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	v201 = v199
	goto L51
L53:
	;
	v191 = v175 + v176<<(uint(int32(3))%32)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v192) {
		goto L52
	} else {
		goto L55
	}
L54:
	;
	v201 = int32(0)
	goto L51
L55:
	;
	v196 = v176 + int32(1)
	if v196 != v170 {
		v176 = v196
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v207 = v201
	goto L50
L58:
	;
	v28 = v207
	v31 = v156
	v32 = v157
	v33 = v158
	goto L1
L59:
	;
	m.G0 = v16 + int32(32)
	return
}
func F_heap_getsysattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v4)
	switch l1 + int32(6) {
	case 0:
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v37 = v19
		m.G0 = v7 + int32(16)
		return v37
	case 1, 3:
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
		v37 = v18
		m.G0 = v7 + int32(16)
		return v37
	case 2:
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
		v37 = v16
		m.G0 = v7 + int32(16)
		return v37
	case 4:
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v37 = v14
		m.G0 = v7 + int32(16)
		return v37
	case 5:
		v37 = l0 + int32(4)
		m.G0 = v7 + int32(16)
		return v37
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
			F_errmsg_internal(m, int32(_a_F_heap_getsysattr_0), v7)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_heap_getsysattr_1), int32(761), int32(_a_F_heap_getsysattr_2))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_heap_prepare_freeze_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v299 int32
	_ = v299
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v478 int32
	_ = v478
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v641 int32
	_ = v641
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v682 int64
	_ = v682
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v721 int32
	_ = v721
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v883 int32
	_ = v883
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v904 int32
	_ = v904
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v922 int32
	_ = v922
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v935 int32
	_ = v935
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v975 int32
	_ = v975
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v989 int32
	_ = v989
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1027 int32
	_ = v1027
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1135 int32
	_ = v1135
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1151 int32
	_ = v1151
	var v1161 int32
	_ = v1161
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1184 int32
	_ = v1184
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1214 int32
	_ = v1214
	var v1220 int32
	_ = v1220
	var v1238 int32
	_ = v1238
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1279 int32
	_ = v1279
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	v6 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(176)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v28
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)) = uint16(v30)
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+8)) = uint16(v6)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v32)
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	v38 = int32(768)
	if v37&v38 != v38 {
		goto L12
	} else {
		goto L13
	}
L1:
	;
	if v83 != 0 {
		goto L338
	} else {
		goto L339
	}
L2:
	;
	v1238 = v1214
	v1244 = v1220
	v1245 = int32(0)
	goto L1
L3:
	;
	v975 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v975)
	if v961&int32(4) != 0 {
		goto L283
	} else {
		goto L284
	}
L4:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v26)+168))
	F_pfree(m, v928)
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L43
	} else {
		goto L270
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L43
	} else {
		goto L266
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L43
	} else {
		goto L262
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L43
	} else {
		goto L257
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L43
	} else {
		goto L253
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L43
	} else {
		goto L249
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L43
	} else {
		goto L245
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L43
	} else {
		goto L241
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = base.B2i32(base.Ui32(v42) < base.Ui32(int32(3)))
	if base.Ui32(v42) < base.Ui32(int32(3)) {
		v78 = v6
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v80 = v37
	v82 = int32(1)
	v83 = v6
	goto L14
L14:
	;
	if base.Ui32(v80&int32(_a_F_heap_prepare_freeze_tuple_0)) < base.Ui32(int32(_a_F_heap_prepare_freeze_tuple_1)) {
		v95 = v80
		v96 = v6
		goto L27
	} else {
		goto L28
	}
L15:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	v80 = v79
	v82 = v44
	v83 = v78
	goto L14
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v45))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v42)) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v57 != 0 {
		goto L11
	} else {
		goto L21
	}
L18:
	;
	v57 = base.B2i32(base.Ui32(v42) < base.Ui32(v45))
	goto L17
L19:
	;
	goto L20
L20:
	;
	v57 = int32(base.Ui32(v42-v45) >> (uint(int32(31)) % 32))
	goto L17
L21:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v58))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v42)) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v70 == int32(0) {
		v78 = v6
		goto L15
	} else {
		goto L26
	}
L23:
	;
	v70 = base.B2i32(base.Ui32(v42) < base.Ui32(v58))
	goto L22
L24:
	;
	goto L25
L25:
	;
	v70 = int32(base.Ui32(v42-v58) >> (uint(int32(31)) % 32))
	goto L22
L26:
	;
	v73 = int32(1)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)))
	v76 = v74 | v73
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v76)
	v78 = v73
	goto L15
L27:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v95&int32(_a_F_heap_prepare_freeze_tuple_2) != 0 {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v88) < base.Ui32(int32(3)) {
		v95 = v80
		v96 = v6
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v91 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v91)
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	v95 = v94
	v96 = v91
	goto L27
L30:
	;
	if v97 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L31:
	;
	v746 = int32(1)
	v747 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	if v747&int32(128)|base.B2i32(v747&int32(_a_F_heap_prepare_freeze_tuple_3) == int32(64)) != 0 {
		v1214 = v746
		v1220 = v6
		goto L2
	} else {
		goto L233
	}
L32:
	;
	v1238 = int32(0)
	v1244 = v6
	v1245 = v6
	goto L1
L33:
	;
	v100 = int32(2)
	if base.B2i32(v97 == int32(0))|base.B2i32(v95&int32(_a_F_heap_prepare_freeze_tuple_4) == int32(_a_F_heap_prepare_freeze_tuple_5)) != 0 {
		v961 = v100
		v962 = v6
		goto L3
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if base.Ui32(v97) < base.Ui32(int32(3)) {
		goto L30
	} else {
		goto L222
	}
L36:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	goto L37
L37:
	;
	if int32(base.Ui32(v97-v108)>>(uint(int32(31))%32)) != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	v113 = v95 & int32(128)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	goto L39
L39:
	;
	if int32(base.Ui32(v97-v114)>>(uint(int32(31))%32)) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v120 = F_MultiXactIdIsRunning(m, v97, base.B2i32(v113 != int32(0)))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v253 = F_GetMultiXactIdMembers(m, v97, v26+int32(168), base.B2i32(v113 != int32(0)))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L43
	} else {
		goto L78
	}
L43:
	;
	return int32(0)
L44:
	;
	if v120 != 0 {
		goto L9
	} else {
		goto L45
	}
L45:
	;
	if v113 != 0 {
		v961 = v100
		v962 = v6
		goto L3
	} else {
		goto L46
	}
L46:
	;
	v127 = F_GetMultiXactIdMembers(m, v97, v26+int32(172), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	if int32(0) < v127 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v26)+172))
	v138 = int32(0)
	goto L53
L49:
	;
	v181 = v6
	goto L50
L50:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v194))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v181)) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L51:
	;
	F_pfree(m, v132)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L43
	} else {
		goto L57
	}
L52:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v168 = v166
	goto L51
L53:
	;
	v158 = v132 + v138<<(uint(int32(3))%32)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v159) {
		goto L52
	} else {
		goto L55
	}
L54:
	;
	v168 = int32(0)
	goto L51
L55:
	;
	v163 = v138 + int32(1)
	if v163 != v127 {
		v138 = v163
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v181 = v168
	goto L50
L58:
	;
	if v206 != 0 {
		goto L8
	} else {
		goto L62
	}
L59:
	;
	v206 = base.B2i32(base.Ui32(v181) < base.Ui32(v194))
	goto L58
L60:
	;
	goto L61
L61:
	;
	v206 = int32(base.Ui32(v181-v194) >> (uint(int32(31)) % 32))
	goto L58
L62:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v207))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v181)) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v219 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v219 = base.B2i32(base.Ui32(v181) < base.Ui32(v207))
	goto L63
L65:
	;
	goto L66
L66:
	;
	v219 = int32(base.Ui32(v181-v207) >> (uint(int32(31)) % 32))
	goto L63
L67:
	;
	v961 = int32(4)
	v962 = v181
	goto L3
L68:
	;
	goto L69
L69:
	;
	v223 = F_TransactionIdDidCommit(m, v181)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L43
	} else {
		goto L70
	}
L70:
	;
	if v223 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v961 = v100
	v962 = int32(0)
	goto L3
L72:
	;
	goto L73
L73:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L43
	} else {
		goto L74
	}
L74:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L43
	} else {
		goto L75
	}
L75:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+88)) = v235
	*(*int32)(unsafe.Add(mBase, uint32(v26)+84)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = v97
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_6), v26+int32(80))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L43
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_7), int32(_a_F_heap_prepare_freeze_tuple_8), int32(_a_F_heap_prepare_freeze_tuple_9))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L43
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	if v253 <= int32(0) {
		v961 = v100
		v962 = v6
		goto L3
	} else {
		goto L79
	}
L79:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v264 = int32(0)
	v268 = v257
	goto L81
L80:
	;
	v333 = F_palloc(m, v253<<(uint(int32(3))%32))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L43
	} else {
		goto L103
	}
L81:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v26)+168))
	v283 = int32(3)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v282+v264<<(uint(v283)%32))))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v287))&base.B2i32(base.Ui32(v283) <= base.Ui32(v286)) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	goto L96
L83:
	;
	if v299 != 0 {
		goto L80
	} else {
		goto L87
	}
L84:
	;
	v299 = base.B2i32(base.Ui32(v286) < base.Ui32(v287))
	goto L83
L85:
	;
	goto L86
L86:
	;
	v299 = int32(base.Ui32(v286-v287) >> (uint(int32(31)) % 32))
	goto L83
L87:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v268))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v286)) == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	if v311 != 0 {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v311 = base.B2i32(base.Ui32(v286) < base.Ui32(v268))
	goto L88
L90:
	;
	goto L91
L91:
	;
	v311 = int32(base.Ui32(v286-v268) >> (uint(int32(31)) % 32))
	goto L88
L92:
	;
	v312 = v286
	goto L94
L93:
	;
	v312 = v268
	goto L94
L94:
	;
	v314 = v264 + int32(1)
	if v314 != v253 {
		v264 = v314
		v268 = v312
		goto L81
	} else {
		goto L95
	}
L95:
	;
	goto L82
L96:
	;
	if int32(base.Ui32(v97-v316)>>(uint(int32(31))%32)) != 0 {
		goto L80
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v312
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L98
L98:
	;
	if int32(base.Ui32(v97-v321)>>(uint(int32(31))%32)) != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v97
	goto L101
L100:
	;
	goto L101
L101:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v26)+168))
	F_pfree(m, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L43
	} else {
		goto L102
	}
L102:
	;
	goto L32
L103:
	;
	v345 = int32(0)
	v346 = v6
	v347 = v6
	v348 = v6
	v350 = v6
	goto L104
L104:
	;
	v359 = int32(3)
	v360 = v345 << (uint(v359) % 32)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v26)+168))
	v362 = v360 + v361
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	if base.Ui32(v364) <= base.Ui32(v359) {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	goto L4
L106:
	;
	v692 = v345 + int32(1)
	if v253 != v692 {
		v345 = v692
		v346 = v687
		v347 = v688
		v348 = v689
		v350 = v690
		goto L104
	} else {
		goto L221
	}
L107:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v26)+168))
	v682 = *(*int64)(unsafe.Add(mBase, uint32(v680+v360)))
	*(*int64)(unsafe.Add(mBase, uint32(v333+v347<<(uint(int32(3))%32)))) = v682
	v687 = v674
	v688 = v347 + int32(1)
	v689 = v675
	v690 = v676
	goto L106
L108:
	;
	if base.Ui32(v363) < base.Ui32(int32(3)) {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	goto L110
L110:
	;
	if v346 != 0 {
		goto L7
	} else {
		goto L165
	}
L111:
	;
	if v486 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L112:
	;
	v486 = int32(0)
	goto L111
L113:
	;
	goto L114
L114:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_freeze_tuple[0]))
	if v377 == v363 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v486 = int32(1)
	goto L111
L116:
	;
	goto L117
L117:
	;
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_freeze_tuple[1]))
	if v381 <= int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v486 = v478
	goto L111
L119:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_freeze_tuple[2]))
	if v385 == int32(0) {
		v478 = int32(0)
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v447 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_freeze_tuple[3]))
	v449 = int32(0)
	v451 = v381 - int32(1)
	goto L141
L122:
	;
	v390 = v385
	goto L123
L123:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v390)+20))
	if v395 == int32(4) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v478 = int32(0)
	goto L118
L125:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v390)+80))
	if v442 != 0 {
		v390 = v442
		goto L123
	} else {
		goto L140
	}
L126:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	if v398 == int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v401 = int32(1)
	if v363 == v398 {
		v478 = v401
		goto L118
	} else {
		goto L128
	}
L128:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v390)+52))
	v405 = v403 - int32(1)
	if v405 < int32(0) {
		goto L125
	} else {
		goto L129
	}
L129:
	;
	v410 = int32(0)
	v412 = v405
	goto L130
L130:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v390)+48))
	v418 = int32(2)
	v419 = base.I32_div_s(v412-v410, v418)
	v420 = v419 + v410
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v416+v420<<(uint(v418)%32))))
	if v424 == v363 {
		v478 = v401
		goto L118
	} else {
		goto L132
	}
L131:
	;
	goto L125
L132:
	;
	v428 = F_TransactionIdPrecedes(m, v424, v363)
	mBase = m.M
	if v428 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v429 = v420 + int32(1)
	goto L135
L134:
	;
	v429 = v410
	goto L135
L135:
	;
	if v428 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v432 = v412
	goto L138
L137:
	;
	v432 = v420 - int32(1)
	goto L138
L138:
	;
	if v429 <= v432 {
		v410 = v429
		v412 = v432
		goto L130
	} else {
		goto L139
	}
L139:
	;
	goto L131
L140:
	;
	goto L124
L141:
	;
	v456 = int32(2)
	v457 = base.I32_div_s(v451-v449, v456)
	v458 = v457 + v449
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v447+v458<<(uint(v456)%32))))
	v463 = base.B2i32(v462 == v363)
	if v462 == v363 {
		v478 = v463
		goto L118
	} else {
		goto L143
	}
L142:
	;
	v478 = v463
	goto L118
L143:
	;
	v466 = base.B2i32(base.Ui32(v462) < base.Ui32(v363))
	if base.Ui32(v462) < base.Ui32(v363) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v467 = v458 + int32(1)
	goto L146
L145:
	;
	v467 = v449
	goto L146
L146:
	;
	if base.Ui32(v462) < base.Ui32(v363) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v470 = v451
	goto L149
L148:
	;
	v470 = v458 - int32(1)
	goto L149
L149:
	;
	if v467 <= v470 {
		v449 = v467
		v451 = v470
		goto L141
	} else {
		goto L150
	}
L150:
	;
	goto L142
L151:
	;
	v489 = F_TransactionIdIsInProgress(m, v363)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L43
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v494))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v363)) == int32(0) {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	if v489 == int32(0) {
		v687 = v346
		v688 = v347
		v689 = v348
		v690 = v350
		goto L106
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	if v506 == int32(0) {
		v674 = v346
		v675 = v348
		v676 = int32(1)
		goto L107
	} else {
		goto L160
	}
L157:
	;
	v506 = base.B2i32(base.Ui32(v363) < base.Ui32(v494))
	goto L156
L158:
	;
	goto L159
L159:
	;
	v506 = int32(base.Ui32(v363-v494) >> (uint(int32(31)) % 32))
	goto L156
L160:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L43
	} else {
		goto L161
	}
L161:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L43
	} else {
		goto L162
	}
L162:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+152)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v26)+148)) = v363
	*(*int32)(unsafe.Add(mBase, uint32(v26)+144)) = v97
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_10), v26+int32(144))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L43
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_7), int32(_a_F_heap_prepare_freeze_tuple_11), int32(_a_F_heap_prepare_freeze_tuple_9))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L43
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	if base.Ui32(v363) < base.Ui32(int32(3)) {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v660))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v363)) == int32(0) {
		goto L217
	} else {
		goto L218
	}
L167:
	;
	if v649 != 0 {
		goto L207
	} else {
		goto L208
	}
L168:
	;
	v649 = int32(0)
	goto L167
L169:
	;
	goto L170
L170:
	;
	v540 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_freeze_tuple[0]))
	if v540 == v363 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v649 = int32(1)
	goto L167
L172:
	;
	goto L173
L173:
	;
	v544 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_freeze_tuple[1]))
	if v544 <= int32(0) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v649 = v641
	goto L167
L175:
	;
	v548 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_freeze_tuple[2]))
	if v548 == int32(0) {
		v641 = int32(0)
		goto L174
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v610 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_freeze_tuple[3]))
	v612 = int32(0)
	v614 = v544 - int32(1)
	goto L197
L178:
	;
	v553 = v548
	goto L179
L179:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v553)+20))
	if v558 == int32(4) {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	v641 = int32(0)
	goto L174
L181:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v553)+80))
	if v605 != 0 {
		v553 = v605
		goto L179
	} else {
		goto L196
	}
L182:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
	if v561 == int32(0) {
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v564 = int32(1)
	if v363 == v561 {
		v641 = v564
		goto L174
	} else {
		goto L184
	}
L184:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v553)+52))
	v568 = v566 - int32(1)
	if v568 < int32(0) {
		goto L181
	} else {
		goto L185
	}
L185:
	;
	v573 = int32(0)
	v575 = v568
	goto L186
L186:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v553)+48))
	v581 = int32(2)
	v582 = base.I32_div_s(v575-v573, v581)
	v583 = v582 + v573
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v579+v583<<(uint(v581)%32))))
	if v587 == v363 {
		v641 = v564
		goto L174
	} else {
		goto L188
	}
L187:
	;
	goto L181
L188:
	;
	v591 = F_TransactionIdPrecedes(m, v587, v363)
	mBase = m.M
	if v591 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v592 = v583 + int32(1)
	goto L191
L190:
	;
	v592 = v573
	goto L191
L191:
	;
	if v591 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v595 = v575
	goto L194
L193:
	;
	v595 = v583 - int32(1)
	goto L194
L194:
	;
	if v592 <= v595 {
		v573 = v592
		v575 = v595
		goto L186
	} else {
		goto L195
	}
L195:
	;
	goto L187
L196:
	;
	goto L180
L197:
	;
	v619 = int32(2)
	v620 = base.I32_div_s(v614-v612, v619)
	v621 = v620 + v612
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v610+v621<<(uint(v619)%32))))
	v626 = base.B2i32(v625 == v363)
	if v625 == v363 {
		v641 = v626
		goto L174
	} else {
		goto L199
	}
L198:
	;
	v641 = v626
	goto L174
L199:
	;
	v629 = base.B2i32(base.Ui32(v625) < base.Ui32(v363))
	if base.Ui32(v625) < base.Ui32(v363) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v630 = v621 + int32(1)
	goto L202
L201:
	;
	v630 = v612
	goto L202
L202:
	;
	if base.Ui32(v625) < base.Ui32(v363) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v633 = v614
	goto L205
L204:
	;
	v633 = v621 - int32(1)
	goto L205
L205:
	;
	if v630 <= v633 {
		v612 = v630
		v614 = v633
		goto L197
	} else {
		goto L206
	}
L206:
	;
	goto L198
L207:
	;
	v658 = v348
	goto L166
L208:
	;
	goto L209
L209:
	;
	v650 = F_TransactionIdIsInProgress(m, v363)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L43
	} else {
		goto L210
	}
L210:
	;
	if v650 != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v658 = v348
	goto L166
L212:
	;
	goto L213
L213:
	;
	v654 = F_TransactionIdDidCommit(m, v363)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L43
	} else {
		goto L214
	}
L214:
	;
	if v654 == int32(0) {
		v687 = int32(0)
		v688 = v347
		v689 = v348
		v690 = v350
		goto L106
	} else {
		goto L215
	}
L215:
	;
	v658 = int32(1)
	goto L166
L216:
	;
	if v672 != 0 {
		goto L6
	} else {
		goto L220
	}
L217:
	;
	v672 = base.B2i32(base.Ui32(v363) < base.Ui32(v660))
	goto L216
L218:
	;
	goto L219
L219:
	;
	v672 = int32(base.Ui32(v363-v660) >> (uint(int32(31)) % 32))
	goto L216
L220:
	;
	v674 = v363
	v675 = v658
	v676 = v350
	goto L107
L221:
	;
	goto L105
L222:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v696))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v97)) == int32(0) {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	if v708 != 0 {
		goto L5
	} else {
		goto L227
	}
L224:
	;
	v708 = base.B2i32(base.Ui32(v97) < base.Ui32(v696))
	goto L223
L225:
	;
	goto L226
L226:
	;
	v708 = int32(base.Ui32(v97-v696) >> (uint(int32(31)) % 32))
	goto L223
L227:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v709))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v97)) == int32(0) {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	if v721 != 0 {
		goto L31
	} else {
		goto L232
	}
L229:
	;
	v721 = base.B2i32(base.Ui32(v97) < base.Ui32(v709))
	goto L228
L230:
	;
	goto L231
L231:
	;
	v721 = int32(base.Ui32(v97-v709) >> (uint(int32(31)) % 32))
	goto L228
L232:
	;
	goto L32
L233:
	;
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)))
	v757 = v755 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v757)
	v1214 = v746
	v1220 = v6
	goto L2
L234:
	;
	v1238 = int32(0)
	v1244 = int32(1)
	v1245 = v6
	goto L1
L235:
	;
	goto L236
L236:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L43
	} else {
		goto L237
	}
L237:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L43
	} else {
		goto L238
	}
L238:
	;
	v770 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v770
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v97
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_12), v26+int32(16))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L43
	} else {
		goto L239
	}
L239:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_7), int32(_a_F_heap_prepare_freeze_tuple_13), int32(_a_F_heap_prepare_freeze_tuple_14))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L43
	} else {
		goto L240
	}
L240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L241:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L43
	} else {
		goto L242
	}
L242:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+164)) = v790
	*(*int32)(unsafe.Add(mBase, uint32(v26)+160)) = v42
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_15), v26+int32(160))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L43
	} else {
		goto L243
	}
L243:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_7), int32(_a_F_heap_prepare_freeze_tuple_16), int32(_a_F_heap_prepare_freeze_tuple_14))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L43
	} else {
		goto L244
	}
L244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L245:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L43
	} else {
		goto L246
	}
L246:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v810
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v97
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_17), v26+int32(32))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L43
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_7), int32(_a_F_heap_prepare_freeze_tuple_18), int32(_a_F_heap_prepare_freeze_tuple_9))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L43
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L249:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L43
	} else {
		goto L250
	}
L250:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v830
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v97
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_19), v26+int32(48))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L43
	} else {
		goto L251
	}
L251:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_7), int32(_a_F_heap_prepare_freeze_tuple_20), int32(_a_F_heap_prepare_freeze_tuple_9))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L43
	} else {
		goto L252
	}
L252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L253:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L43
	} else {
		goto L254
	}
L254:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = v850
	*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v97
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_21), v26-int32(-64))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L43
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_7), int32(_a_F_heap_prepare_freeze_tuple_22), int32(_a_F_heap_prepare_freeze_tuple_9))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L43
	} else {
		goto L256
	}
L256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L257:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L43
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+128)) = v97
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_23), v26+int32(128))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L43
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+116)) = v363
	*(*int32)(unsafe.Add(mBase, uint32(v26)+112)) = v346
	F_errdetail_internal(m, int32(_a_F_heap_prepare_freeze_tuple_24), v26+int32(112))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L43
	} else {
		goto L260
	}
L260:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_7), int32(_a_F_heap_prepare_freeze_tuple_25), int32(_a_F_heap_prepare_freeze_tuple_9))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L43
	} else {
		goto L261
	}
L261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L262:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L43
	} else {
		goto L263
	}
L263:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+104)) = v896
	*(*int32)(unsafe.Add(mBase, uint32(v26)+100)) = v363
	*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v97
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_6), v26+int32(96))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L43
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_7), int32(_a_F_heap_prepare_freeze_tuple_26), int32(_a_F_heap_prepare_freeze_tuple_9))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L43
	} else {
		goto L265
	}
L265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L266:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L43
	} else {
		goto L267
	}
L267:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v917
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v97
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_27), v26)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L43
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_7), int32(_a_F_heap_prepare_freeze_tuple_28), int32(_a_F_heap_prepare_freeze_tuple_14))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L43
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L270:
	;
	if v688 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	F_pfree(m, v333)
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L43
	} else {
		goto L282
	}
L272:
	;
	v948 = int32(2)
	v949 = int32(0)
	goto L271
L273:
	;
	goto L274
L274:
	;
	v935 = int32(0)
	if v690|base.B2i32(v687 == v935) == v935 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	if v689&int32(1) != 0 {
		goto L278
	} else {
		goto L279
	}
L276:
	;
	goto L277
L277:
	;
	v946 = F_MultiXactIdCreateFromMembers(m, v688, v333)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L43
	} else {
		goto L281
	}
L278:
	;
	v944 = int32(20)
	goto L280
L279:
	;
	v944 = int32(4)
	goto L280
L280:
	;
	v948 = v944
	v949 = v687
	goto L271
L281:
	;
	v948 = int32(8)
	v949 = v946
	goto L271
L282:
	;
	v961 = v948
	v962 = v949
	goto L3
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v962
	v981 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
	v983 = v981 & int32(-7377)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v983)
	if v961&int32(16) != 0 {
		goto L286
	} else {
		goto L287
	}
L284:
	;
	goto L285
L285:
	;
	if v961&int32(8) == int32(0) {
		goto L289
	} else {
		goto L290
	}
L286:
	;
	v989 = v983 | int32(1024)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v989)
	goto L288
L287:
	;
	goto L288
L288:
	;
	v1238 = int32(0)
	v1244 = int32(0)
	v1245 = v975
	goto L1
L289:
	;
	v1214 = int32(1)
	v1220 = int32(0)
	goto L2
L290:
	;
	goto L291
L291:
	;
	v998 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
	v1000 = v998 & int32(_a_F_heap_prepare_freeze_tuple_29)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v1000)
	v1002 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	v1004 = v1002 & int32(_a_F_heap_prepare_freeze_tuple_30)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)) = uint16(v1004)
	v1006 = int32(0)
	v1010 = F_GetMultiXactIdMembers(m, v962, v26+int32(172), v1006)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L43
	} else {
		goto L293
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v962
	v1202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
	v1203 = v1202 | v1200
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v1203)
	v1205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	v1206 = v1205 | v1184
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)) = uint16(v1206)
	v1238 = int32(0)
	v1244 = v1006
	v1245 = v975
	goto L1
L293:
	;
	if v1010 <= int32(0) {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1184 = int32(0)
	v1200 = int32(_a_F_heap_prepare_freeze_tuple_31)
	goto L292
L295:
	;
	goto L296
L296:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v26)+172))
	if v1010 == int32(1) {
		goto L299
	} else {
		goto L300
	}
L297:
	;
	F_pfree(m, v1016)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L43
	} else {
		goto L326
	}
L298:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1016+v1106<<(uint(int32(3))%32))+4))
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1123<<(uint(int32(2))%32))+uint32(_c_F_heap_prepare_freeze_tuple[4])))
	if base.Ui32(v1102) < base.Ui32(v1126) {
		goto L320
	} else {
		goto L321
	}
L299:
	;
	v1019 = int32(0)
	v1102 = v1019
	v1104 = v1019
	v1106 = v1019
	v1111 = v1019
	goto L298
L300:
	;
	goto L301
L301:
	;
	v1027 = int32(0)
	v1037 = v1027
	v1039 = v1027
	v1041 = v1027
	v1045 = v1027
	v1046 = v1027
	goto L302
L302:
	;
	v1057 = v1016 + v1041<<(uint(int32(3))%32)
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+4))
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1058<<(uint(int32(2))%32))+uint32(_c_F_heap_prepare_freeze_tuple[4])))
	if base.Ui32(v1037) < base.Ui32(v1061) {
		goto L304
	} else {
		goto L305
	}
L303:
	;
	if v1010&int32(1) == int32(0) {
		v1142 = v1089
		v1144 = v1087
		v1151 = v1088
		goto L297
	} else {
		goto L319
	}
L304:
	;
	v1063 = v1061
	goto L306
L305:
	;
	v1063 = v1037
	goto L306
L306:
	;
	switch v1058 - int32(3) {
	case 0:
		goto L310
	case 1:
		v1070 = v1039
		goto L308
	case 2:
		goto L309
	default:
		v1072 = v1039
		v1073 = v1046
		goto L307
	}
L307:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+12))
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1074<<(uint(int32(2))%32))+uint32(_c_F_heap_prepare_freeze_tuple[4])))
	switch v1074 - int32(3) {
	case 0:
		goto L314
	case 1:
		v1085 = v1072
		goto L312
	case 2:
		goto L313
	default:
		v1087 = v1072
		v1088 = v1073
		goto L311
	}
L308:
	;
	v1072 = v1070
	v1073 = int32(1)
	goto L307
L309:
	;
	v1070 = v1039 | int32(_a_F_heap_prepare_freeze_tuple_32)
	goto L308
L310:
	;
	v1072 = v1039 | int32(_a_F_heap_prepare_freeze_tuple_32)
	v1073 = v1046
	goto L307
L311:
	;
	if base.Ui32(v1063) < base.Ui32(v1077) {
		goto L315
	} else {
		goto L316
	}
L312:
	;
	v1087 = v1085
	v1088 = int32(1)
	goto L311
L313:
	;
	v1085 = v1072 | int32(_a_F_heap_prepare_freeze_tuple_32)
	goto L312
L314:
	;
	v1087 = v1072 | int32(_a_F_heap_prepare_freeze_tuple_32)
	v1088 = v1073
	goto L311
L315:
	;
	v1089 = v1077
	goto L317
L316:
	;
	v1089 = v1063
	goto L317
L317:
	;
	v1090 = int32(2)
	v1091 = v1041 + v1090
	v1093 = v1045 + v1090
	if v1093 != v1010&int32(2147483646) {
		v1037 = v1089
		v1039 = v1087
		v1041 = v1091
		v1045 = v1093
		v1046 = v1088
		goto L302
	} else {
		goto L318
	}
L318:
	;
	goto L303
L319:
	;
	v1102 = v1089
	v1104 = v1087
	v1106 = v1091
	v1111 = v1088
	goto L298
L320:
	;
	v1128 = v1126
	goto L322
L321:
	;
	v1128 = v1102
	goto L322
L322:
	;
	switch v1123 - int32(3) {
	case 0:
		goto L325
	case 1:
		v1135 = v1104
		goto L323
	case 2:
		goto L324
	default:
		v1142 = v1128
		v1144 = v1104
		v1151 = v1111
		goto L297
	}
L323:
	;
	v1142 = v1128
	v1144 = v1135
	v1151 = int32(1)
	goto L297
L324:
	;
	v1135 = v1104 | int32(_a_F_heap_prepare_freeze_tuple_32)
	goto L323
L325:
	;
	v1142 = v1128
	v1144 = v1104 | int32(_a_F_heap_prepare_freeze_tuple_32)
	v1151 = v1111
	goto L297
L326:
	;
	if v1142&int32(-2) == int32(2) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	if v1151 != 0 {
		v1184 = v1144
		v1200 = int32(_a_F_heap_prepare_freeze_tuple_33)
		goto L292
	} else {
		goto L330
	}
L328:
	;
	goto L329
L329:
	;
	if v1142 != 0 {
		goto L331
	} else {
		goto L332
	}
L330:
	;
	v1184 = v1144
	v1200 = int32(_a_F_heap_prepare_freeze_tuple_34)
	goto L292
L331:
	;
	v1171 = int32(_a_F_heap_prepare_freeze_tuple_2)
	goto L333
L332:
	;
	v1171 = int32(_a_F_heap_prepare_freeze_tuple_35)
	goto L333
L333:
	;
	if v1142 == int32(1) {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v1174 = int32(_a_F_heap_prepare_freeze_tuple_3)
	goto L336
L335:
	;
	v1174 = v1171
	goto L336
L336:
	;
	if v1151 != 0 {
		v1184 = v1144
		v1200 = v1174
		goto L292
	} else {
		goto L337
	}
L337:
	;
	v1184 = v1144
	v1200 = v1174 | int32(128)
	goto L292
L338:
	;
	v1256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
	v1258 = v1256 | int32(768)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v1258)
	goto L340
L339:
	;
	goto L340
L340:
	;
	if v96 != 0 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v1260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)))
	v1263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	if v1263&int32(_a_F_heap_prepare_freeze_tuple_1) != 0 {
		goto L344
	} else {
		goto L345
	}
L342:
	;
	goto L343
L343:
	;
	if v1238 != 0 {
		goto L347
	} else {
		goto L348
	}
L344:
	;
	v1266 = int32(4)
	goto L346
L345:
	;
	v1266 = int32(2)
	goto L346
L346:
	;
	v1267 = v1260 | v1266
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v1267)
	goto L343
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	v1271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	v1273 = v1271 & int32(_a_F_heap_prepare_freeze_tuple_36)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)) = uint16(v1273)
	v1275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
	v1279 = v1275&int32(_a_F_heap_prepare_freeze_tuple_29) | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v1279)
	goto L349
L348:
	;
	goto L349
L349:
	;
	v1283 = (v82 | v83) & (v1238 | v1244)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v1283)
	if v1244&v82 != 0 {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	m.G0 = v26 + int32(176)
	return v83 | (v1245 | v96) | v1238
L351:
	;
	v1286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v1286&int32(1) != 0 {
		goto L350
	} else {
		goto L352
	}
L352:
	;
	v1293 = F_heap_tuple_should_freeze(m, l0, l1, l2+int32(12), l2+int32(16))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L43
	} else {
		goto L353
	}
L353:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v1293)
	goto L350
}
func F_heap_prepare_pagescan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v218 int32
	_ = v218
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v285 int32
	_ = v285
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v352 int32
	_ = v352
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v410 int32
	_ = v410
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_heap_page_prune_opt(m, v21, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_LockBuffer(m, v22, int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v22 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v46) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_pagescan[0]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31+(v22^int32(-1))<<(uint(int32(2))%32))))
	v45 = v37
	goto L4
L6:
	;
	goto L7
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_pagescan[1]))
	v45 = v39 + v22<<(uint(int32(13))%32) + int32(-8192)
	goto L4
L8:
	;
	v56 = int32(base.Ui32(v46+int32(_a_F_heap_prepare_pagescan_0))>>(uint(int32(2))%32)) & int32(_a_F_heap_prepare_pagescan_1)
	goto L10
L9:
	;
	v56 = int32(0)
	goto L10
L10:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+10)))
	if v57&int32(4) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v394
	F_LockBuffer(m, v22, int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L77
	}
L12:
	;
	if v56 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L13:
	;
	if v56 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v63 = F_CheckForSerializableConflictOutNeeded(m, v62, v19)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+29)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v69 = F_CheckForSerializableConflictOutNeeded(m, v68, v19)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	if v63 == int32(0) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L12
L19:
	;
	if v67 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v69 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	goto L22
L22:
	;
	if v69 != 0 {
		goto L12
	} else {
		goto L55
	}
L23:
	;
	v195 = int32(base.Ui32(v20) >> (uint(int32(16)) % 32))
	v199 = int32(1)
	v202 = int32(0)
	v203 = v199
	v204 = v199
	goto L48
L24:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v82+v166<<(uint(int32(2))%32))))
	if v181&int32(_a_F_heap_prepare_pagescan_2) != int32(_a_F_heap_prepare_pagescan_3) {
		v394 = v165
		goto L11
	} else {
		goto L47
	}
L25:
	;
	v92 = int32(2)
	if base.Ui32(v86) <= base.Ui32(v92) {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	if v56 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v56 != 0 {
		goto L23
	} else {
		goto L33
	}
L29:
	;
	v394 = int32(0)
	goto L11
L30:
	;
	goto L31
L31:
	;
	v78 = int32(1)
	v80 = l0 + int32(108)
	v82 = v45 + int32(20)
	v86 = (v56 + v78) & int32(_a_F_heap_prepare_pagescan_1)
	if base.Ui32(int32(3)) <= base.Ui32(v86) {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	v165 = int32(0)
	v166 = v78
	v167 = int32(1)
	goto L24
L33:
	;
	v394 = int32(0)
	goto L11
L34:
	;
	v95 = v92
	goto L36
L35:
	;
	v95 = v86
	goto L36
L36:
	;
	v96 = int32(1)
	v97 = v95 - v96
	v102 = int32(0)
	v106 = v102
	v107 = v78
	v108 = v96
	v112 = v102
	goto L37
L37:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v82+v107<<(uint(int32(2))%32))))
	if v122&int32(_a_F_heap_prepare_pagescan_2) == int32(_a_F_heap_prepare_pagescan_3) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v97&v96 == int32(0) {
		v394 = v152
		goto L11
	} else {
		goto L46
	}
L39:
	;
	v127 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v80+v106<<(uint(v127)%32)))) = uint16(v108)
	v133 = v106 + v127
	goto L41
L40:
	;
	v133 = v106
	goto L41
L41:
	;
	v135 = v108 + int32(1)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v82+v135&int32(_a_F_heap_prepare_pagescan_1)<<(uint(int32(2))%32))))
	if v141&int32(_a_F_heap_prepare_pagescan_2) == int32(_a_F_heap_prepare_pagescan_3) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v146 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v80+v133<<(uint(v146)%32)))) = uint16(v135)
	v152 = v133 + v146
	goto L44
L43:
	;
	v152 = v133
	goto L44
L44:
	;
	v153 = int32(2)
	v154 = v108 + v153
	v155 = int32(_a_F_heap_prepare_pagescan_1)
	v156 = v154 & v155
	v158 = v112 + v153
	if v158&v155 != v97&int32(_a_F_heap_prepare_pagescan_4) {
		v106 = v152
		v107 = v156
		v108 = v154
		v112 = v158
		goto L37
	} else {
		goto L45
	}
L45:
	;
	goto L38
L46:
	;
	v165 = v152
	v166 = v156
	v167 = v154
	goto L24
L47:
	;
	v186 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v80+v165<<(uint(v186)%32)))) = uint16(v167)
	v394 = v165 + v186
	goto L11
L48:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(20)+v203<<(uint(int32(2))%32))))
	if v218&int32(_a_F_heap_prepare_pagescan_2) == int32(_a_F_heap_prepare_pagescan_3) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v394 = v247
	goto L11
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(base.Ui32(v218) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v45 + v218&int32(_a_F_heap_prepare_pagescan_5)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+20)) = uint16(v204)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+18)) = uint16(v20)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)) = uint16(v195)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v231
	F_HeapCheckForSerializableConflictOut(m, int32(1), v230, v17+int32(12), v22, v19)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	v247 = v202
	goto L52
L52:
	;
	v251 = v204 + int32(1)
	v253 = v251 & int32(_a_F_heap_prepare_pagescan_1)
	if base.Ui32(v253) <= base.Ui32(v56) {
		v202 = v247
		v203 = v253
		v204 = v251
		goto L48
	} else {
		goto L54
	}
L53:
	;
	v241 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(108)+v202<<(uint(v241)%32)))) = uint16(v204)
	v247 = v202 + v241
	goto L52
L54:
	;
	goto L49
L55:
	;
	goto L13
L56:
	;
	v394 = int32(0)
	goto L11
L57:
	;
	goto L58
L58:
	;
	v262 = int32(base.Ui32(v20) >> (uint(int32(16)) % 32))
	v266 = int32(1)
	v269 = int32(0)
	v270 = v266
	v271 = v266
	goto L59
L59:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(20)+v270<<(uint(int32(2))%32))))
	if v285&int32(_a_F_heap_prepare_pagescan_2) != int32(_a_F_heap_prepare_pagescan_3) {
		v315 = v269
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v394 = v315
	goto L11
L61:
	;
	v318 = v271 + int32(1)
	v320 = v318 & int32(_a_F_heap_prepare_pagescan_1)
	if base.Ui32(v320) <= base.Ui32(v56) {
		v269 = v315
		v270 = v320
		v271 = v318
		goto L59
	} else {
		goto L65
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(base.Ui32(v285) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v45 + v285&int32(_a_F_heap_prepare_pagescan_5)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+20)) = uint16(v271)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+18)) = uint16(v20)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)) = uint16(v262)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v298
	v305 = F_HeapTupleSatisfiesVisibility(m, v17+int32(12), v19, v22)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v305 == int32(0) {
		v315 = v269
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v309 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(108)+v269<<(uint(v309)%32)))) = uint16(v271)
	v315 = v269 + v309
	goto L61
L65:
	;
	goto L60
L66:
	;
	v394 = int32(0)
	goto L11
L67:
	;
	goto L68
L68:
	;
	v329 = int32(base.Ui32(v20) >> (uint(int32(16)) % 32))
	v333 = int32(1)
	v336 = int32(0)
	v337 = v333
	v338 = v333
	goto L69
L69:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(20)+v337<<(uint(int32(2))%32))))
	if v352&int32(_a_F_heap_prepare_pagescan_2) != int32(_a_F_heap_prepare_pagescan_3) {
		v385 = v336
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v394 = v385
	goto L11
L71:
	;
	v389 = v338 + int32(1)
	v391 = v389 & int32(_a_F_heap_prepare_pagescan_1)
	if base.Ui32(v391) <= base.Ui32(v56) {
		v336 = v385
		v337 = v391
		v338 = v389
		goto L69
	} else {
		goto L76
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(base.Ui32(v352) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v45 + v352&int32(_a_F_heap_prepare_pagescan_5)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+20)) = uint16(v338)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+18)) = uint16(v20)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)) = uint16(v329)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v365
	v371 = v17 + int32(12)
	v372 = F_HeapTupleSatisfiesVisibility(m, v371, v19, v22)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_HeapCheckForSerializableConflictOut(m, v372, v374, v371, v22, v19)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if v372 == int32(0) {
		v385 = v336
		goto L71
	} else {
		goto L75
	}
L75:
	;
	v379 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(108)+v336<<(uint(v379)%32)))) = uint16(v338)
	v385 = v336 + v379
	goto L71
L76:
	;
	goto L70
L77:
	;
	m.G0 = v17 + int32(32)
	return
}
func F_heap_rescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	if l2 == int32(0) {
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if l3 != 0 {
			v14 = int32(64)
		} else {
			v14 = int32(0)
		}
		if l4 != 0 {
			v18 = int32(128)
		} else {
			v18 = int32(0)
		}
		v19 = v9&int32(-193) | v14 | v18
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v19
		if l5 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v19 & int32(-257)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v23 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v19 & int32(-257)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				switch v26 {
				case 0, 5:
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v19 | int32(256)
				default:
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v19 & int32(-257)
				}
			}
		}
	}
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v36 != 0 {
		F_ReleaseBuffer(m, v36)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
			if v41 != 0 {
				F_read_stream_reset(m, v41)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					F_initscan(m, l0, l1, int32(1))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				F_initscan(m, l0, l1, int32(1))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		if v41 != 0 {
			F_read_stream_reset(m, v41)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				F_initscan(m, l0, l1, int32(1))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F_initscan(m, l0, l1, int32(1))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				return
			}
		}
	}
}
