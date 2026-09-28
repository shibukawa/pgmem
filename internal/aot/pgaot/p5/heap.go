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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
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
	var v159 int32
	_ = v159
	v4 = int32(0)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	if v5&int32(1) != 0 {
		v159 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v159)
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
	if v140 == int32(0) {
		v159 = v4
		goto L1
	} else {
		goto L43
	}
L4:
	;
	v140 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderAdjustCmax[0]))
	if v20 == v8 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v140 = int32(1)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderAdjustCmax[1]))
	if v24 <= int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v140 = v130
	goto L3
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderAdjustCmax[2]))
	if v28 == int32(0) {
		v130 = int32(0)
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderAdjustCmax[3]))
	v100 = int32(0)
	v103 = v24 - int32(1)
	goto L33
L14:
	;
	v33 = v28
	goto L15
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	if v39 == int32(4) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v130 = int32(0)
	goto L10
L17:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v33)+80))
	if v93 != 0 {
		v33 = v93
		goto L15
	} else {
		goto L32
	}
L18:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v42 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v45 = int32(1)
	if v8 == v42 {
		v130 = v45
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v49 = v47 - int32(1)
	if v49 < int32(0) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v55 = int32(0)
	v58 = v49
	goto L22
L22:
	;
	v63 = int32(2)
	v64 = base.I32_div_s(v58-v55, v63)
	v65 = v64 + v55
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v52+v65<<(uint(v63)%32))))
	if v69 == v8 {
		v130 = v45
		goto L10
	} else {
		goto L24
	}
L23:
	;
	goto L17
L24:
	;
	v78 = base.B2i32(v69-v8 < int32(0)) | base.B2i32(base.Ui32(v69) < base.Ui32(int32(3)))
	if v78 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v79 = v65 + int32(1)
	goto L27
L26:
	;
	v79 = v55
	goto L27
L27:
	;
	if v78 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v82 = v58
	goto L30
L29:
	;
	v82 = v65 - int32(1)
	goto L30
L30:
	;
	if v79 <= v82 {
		v55 = v79
		v58 = v82
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
	v108 = int32(2)
	v109 = base.I32_div_s(v103-v100, v108)
	v110 = v109 + v100
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v98+v110<<(uint(v108)%32))))
	v115 = base.B2i32(v114 == v8)
	if v114 == v8 {
		v130 = v115
		goto L10
	} else {
		goto L35
	}
L34:
	;
	v130 = v115
	goto L10
L35:
	;
	v118 = base.B2i32(base.Ui32(v114) < base.Ui32(v8))
	if base.Ui32(v114) < base.Ui32(v8) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v119 = v110 + int32(1)
	goto L38
L37:
	;
	v119 = v100
	goto L38
L38:
	;
	if base.Ui32(v114) < base.Ui32(v8) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v122 = v103
	goto L41
L40:
	;
	v122 = v110 - int32(1)
	goto L41
L41:
	;
	if v119 <= v122 {
		v100 = v119
		v103 = v122
		goto L33
	} else {
		goto L42
	}
L42:
	;
	goto L34
L43:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v144&int32(32) != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderAdjustCmax[4]))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148+v143<<(uint(int32(3))%32))))
	v153 = v152
	goto L46
L45:
	;
	v153 = v143
	goto L46
L46:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v155 = F_GetComboCommandId(m, v153, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	return
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v155
	v159 = int32(1)
	goto L1
}
func F_heap_abort_speculative(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v122 int32
	_ = v122
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
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int64
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
		v23 = F_ReadBuffer(m, l0, v18|v19<<(uint(int32(16))%32))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			if v23 < int32(0) {
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[0]))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v28+(v23^int32(-1))<<(uint(int32(2))%32))))
				v42 = v34
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[1]))
				v42 = v36 + v23<<(uint(int32(13))%32) + int32(-8192)
			}
			F_LockBufferInternal(m, v23, int32(3))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v47
				v50 = v42 + int32(20)
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v46<<(uint(int32(2))%32))))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = int32(base.Ui32(v54) >> (uint(int32(17)) % 32))
				v60 = v42 + v54&int32(_a_F_heap_abort_speculative_0)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v60
				v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				*(*uint16)(unsafe.Add(mBase, uint32(v14)+20)) = uint16(v62)
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v64
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
				if v16 == v66 {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+68))
					if v69 != int32(99) {
						v72 = F_isTempToastNamespace(m, v69)
						mBase = m.M
						v74 = v72
					} else {
						v74 = int32(1)
					}
					if v74 == int32(0) {
						v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+16)))
						if v77 != int32(_a_F_heap_abort_speculative_1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v237 = m.ExcPending
							if v237 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_heap_abort_speculative_2), int32(0))
								mBase = m.M
								v241 = m.ExcPending
								if v241 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_heap_abort_speculative_3), int32(_a_F_heap_abort_speculative_4), int32(_a_F_heap_abort_speculative_5))
									mBase = m.M
									v246 = m.ExcPending
									if v246 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v80 = int32(_a_F_heap_abort_speculative_6)
							v82 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
							*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v82 + int32(1)
							v87 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[3]))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+136))
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v87))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v91)) == int32(0) {
								v101 = base.B2i32(base.Ui32(v87) < base.Ui32(v91))
							} else {
								v101 = int32(base.Ui32(v87-v91) >> (uint(int32(31)) % 32))
							}
							v103 = v14 + int32(16)
							if v101 != 0 {
								v104 = v91
							} else {
								v104 = v87
							}
							v105 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
							if v105 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v50))) = v104
							} else {
								v108 = int32(3)
								if base.B2i32(base.Ui32(v104) < base.Ui32(v108))|base.B2i32(base.Ui32(v105) < base.Ui32(v108)) == int32(0) {
									if v104-v105 < int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v50))) = v104
									} else {
									}
								} else {
									if base.Ui32(v105) <= base.Ui32(v104) {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v50))) = v104
									}
								}
							}
							*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(0)
							v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+20)))
							v124 = v122 & int32(_a_F_heap_abort_speculative_7)
							*(*uint16)(unsafe.Add(mBase, uint32(v60)+20)) = uint16(v124)
							v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+18)))
							v128 = v126 & int32(_a_F_heap_abort_speculative_8)
							*(*uint16)(unsafe.Add(mBase, uint32(v60)+18)) = uint16(v128)
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
							*(*int32)(unsafe.Add(mBase, uint32(v60)+12)) = v130
							v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+4)))
							*(*uint16)(unsafe.Add(mBase, uint32(v60)+16)) = uint16(v132)
							F_MarkBufferDirty(m, v23)
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return
							} else {
								v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+118)))
								if v137 != int32(112) {
									v198 = int32(_a_F_heap_abort_speculative_6)
									v200 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
									*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v200 - int32(1)
									F_UnlockBuffer(m, v23)
									mBase = m.M
									v205 = m.ExcPending
									if v205 != 0 {
										return
									} else {
										v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+20)))
										if v206&int32(4) != 0 {
											F_heap_toast_delete(m, l0, v14+int32(12), int32(1))
											mBase = m.M
											v213 = m.ExcPending
											if v213 != 0 {
												return
											} else {
												F_ReleaseBuffer(m, v23)
												mBase = m.M
												v215 = m.ExcPending
												if v215 != 0 {
													return
												} else {
													F_pgstat_count_heap_delete(m, l0)
													mBase = m.M
													v217 = m.ExcPending
													if v217 != 0 {
														return
													} else {
														m.G0 = v14 + int32(32)
														return
													}
												}
											}
										} else {
											F_ReleaseBuffer(m, v23)
											mBase = m.M
											v215 = m.ExcPending
											if v215 != 0 {
												return
											} else {
												F_pgstat_count_heap_delete(m, l0)
												mBase = m.M
												v217 = m.ExcPending
												if v217 != 0 {
													return
												} else {
													m.G0 = v14 + int32(32)
													return
												}
											}
										}
									}
								} else {
									v141 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[4]))
									if v141 <= int32(0) {
										v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										if v144 != 0 {
											v198 = int32(_a_F_heap_abort_speculative_6)
											v200 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
											*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v200 - int32(1)
											F_UnlockBuffer(m, v23)
											mBase = m.M
											v205 = m.ExcPending
											if v205 != 0 {
												return
											} else {
												v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+20)))
												if v206&int32(4) != 0 {
													F_heap_toast_delete(m, l0, v14+int32(12), int32(1))
													mBase = m.M
													v213 = m.ExcPending
													if v213 != 0 {
														return
													} else {
														F_ReleaseBuffer(m, v23)
														mBase = m.M
														v215 = m.ExcPending
														if v215 != 0 {
															return
														} else {
															F_pgstat_count_heap_delete(m, l0)
															mBase = m.M
															v217 = m.ExcPending
															if v217 != 0 {
																return
															} else {
																m.G0 = v14 + int32(32)
																return
															}
														}
													}
												} else {
													F_ReleaseBuffer(m, v23)
													mBase = m.M
													v215 = m.ExcPending
													if v215 != 0 {
														return
													} else {
														F_pgstat_count_heap_delete(m, l0)
														mBase = m.M
														v217 = m.ExcPending
														if v217 != 0 {
															return
														} else {
															m.G0 = v14 + int32(32)
															return
														}
													}
												}
											}
										} else {
											v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											if v145 != 0 {
												v198 = int32(_a_F_heap_abort_speculative_6)
												v200 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
												*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v200 - int32(1)
												F_UnlockBuffer(m, v23)
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
													return
												} else {
													v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+20)))
													if v206&int32(4) != 0 {
														F_heap_toast_delete(m, l0, v14+int32(12), int32(1))
														mBase = m.M
														v213 = m.ExcPending
														if v213 != 0 {
															return
														} else {
															F_ReleaseBuffer(m, v23)
															mBase = m.M
															v215 = m.ExcPending
															if v215 != 0 {
																return
															} else {
																F_pgstat_count_heap_delete(m, l0)
																mBase = m.M
																v217 = m.ExcPending
																if v217 != 0 {
																	return
																} else {
																	m.G0 = v14 + int32(32)
																	return
																}
															}
														}
													} else {
														F_ReleaseBuffer(m, v23)
														mBase = m.M
														v215 = m.ExcPending
														if v215 != 0 {
															return
														} else {
															F_pgstat_count_heap_delete(m, l0)
															mBase = m.M
															v217 = m.ExcPending
															if v217 != 0 {
																return
															} else {
																m.G0 = v14 + int32(32)
																return
															}
														}
													}
												}
											} else {
												v146 = int32(8)
												*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)) = uint8(v146)
												v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+18)))
												v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+20)))
												v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+20)))
												*(*uint16)(unsafe.Add(mBase, uint32(v14)+8)) = uint16(v150)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v16
												v157 = int32(1)
												v161 = int32(4)
												v176 = int32(base.Ui32(v148)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v149)>>(uint(v157)%32))&v146 | (int32(base.Ui32(v149)>>(uint(v161)%32))&v161 | (int32(base.Ui32(v149)>>(uint(int32(12))%32))&v157 | int32(base.Ui32(v149)>>(uint(int32(6))%32))&int32(2))))
												*(*uint8)(unsafe.Add(mBase, uint32(v14)+10)) = uint8(v176)
												F_XLogBeginInsert(m)
												mBase = m.M
												v179 = m.ExcPending
												if v179 != 0 {
													return
												} else {
													F_XLogRegisterData(m, v14+int32(4), int32(8))
													mBase = m.M
													v184 = m.ExcPending
													if v184 != 0 {
														return
													} else {
														F_XLogRegisterBuffer(m, int32(0), v23, int32(8))
														mBase = m.M
														v188 = m.ExcPending
														if v188 != 0 {
															return
														} else {
															v191 = F_XLogInsert(m, int32(10), int32(16))
															mBase = m.M
															v192 = m.ExcPending
															if v192 != 0 {
																return
															} else {
																*(*int64)(unsafe.Add(mBase, uint32(v42))) = base.I64_rotl(v191, int64(32))
																v198 = int32(_a_F_heap_abort_speculative_6)
																v200 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
																*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v200 - int32(1)
																F_UnlockBuffer(m, v23)
																mBase = m.M
																v205 = m.ExcPending
																if v205 != 0 {
																	return
																} else {
																	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+20)))
																	if v206&int32(4) != 0 {
																		F_heap_toast_delete(m, l0, v14+int32(12), int32(1))
																		mBase = m.M
																		v213 = m.ExcPending
																		if v213 != 0 {
																			return
																		} else {
																			F_ReleaseBuffer(m, v23)
																			mBase = m.M
																			v215 = m.ExcPending
																			if v215 != 0 {
																				return
																			} else {
																				F_pgstat_count_heap_delete(m, l0)
																				mBase = m.M
																				v217 = m.ExcPending
																				if v217 != 0 {
																					return
																				} else {
																					m.G0 = v14 + int32(32)
																					return
																				}
																			}
																		}
																	} else {
																		F_ReleaseBuffer(m, v23)
																		mBase = m.M
																		v215 = m.ExcPending
																		if v215 != 0 {
																			return
																		} else {
																			F_pgstat_count_heap_delete(m, l0)
																			mBase = m.M
																			v217 = m.ExcPending
																			if v217 != 0 {
																				return
																			} else {
																				m.G0 = v14 + int32(32)
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
										v146 = int32(8)
										*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)) = uint8(v146)
										v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+18)))
										v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+20)))
										v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+20)))
										*(*uint16)(unsafe.Add(mBase, uint32(v14)+8)) = uint16(v150)
										*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v16
										v157 = int32(1)
										v161 = int32(4)
										v176 = int32(base.Ui32(v148)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v149)>>(uint(v157)%32))&v146 | (int32(base.Ui32(v149)>>(uint(v161)%32))&v161 | (int32(base.Ui32(v149)>>(uint(int32(12))%32))&v157 | int32(base.Ui32(v149)>>(uint(int32(6))%32))&int32(2))))
										*(*uint8)(unsafe.Add(mBase, uint32(v14)+10)) = uint8(v176)
										F_XLogBeginInsert(m)
										mBase = m.M
										v179 = m.ExcPending
										if v179 != 0 {
											return
										} else {
											F_XLogRegisterData(m, v14+int32(4), int32(8))
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
												return
											} else {
												F_XLogRegisterBuffer(m, int32(0), v23, int32(8))
												mBase = m.M
												v188 = m.ExcPending
												if v188 != 0 {
													return
												} else {
													v191 = F_XLogInsert(m, int32(10), int32(16))
													mBase = m.M
													v192 = m.ExcPending
													if v192 != 0 {
														return
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v42))) = base.I64_rotl(v191, int64(32))
														v198 = int32(_a_F_heap_abort_speculative_6)
														v200 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
														*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v200 - int32(1)
														F_UnlockBuffer(m, v23)
														mBase = m.M
														v205 = m.ExcPending
														if v205 != 0 {
															return
														} else {
															v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+20)))
															if v206&int32(4) != 0 {
																F_heap_toast_delete(m, l0, v14+int32(12), int32(1))
																mBase = m.M
																v213 = m.ExcPending
																if v213 != 0 {
																	return
																} else {
																	F_ReleaseBuffer(m, v23)
																	mBase = m.M
																	v215 = m.ExcPending
																	if v215 != 0 {
																		return
																	} else {
																		F_pgstat_count_heap_delete(m, l0)
																		mBase = m.M
																		v217 = m.ExcPending
																		if v217 != 0 {
																			return
																		} else {
																			m.G0 = v14 + int32(32)
																			return
																		}
																	}
																}
															} else {
																F_ReleaseBuffer(m, v23)
																mBase = m.M
																v215 = m.ExcPending
																if v215 != 0 {
																	return
																} else {
																	F_pgstat_count_heap_delete(m, l0)
																	mBase = m.M
																	v217 = m.ExcPending
																	if v217 != 0 {
																		return
																	} else {
																		m.G0 = v14 + int32(32)
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
						v80 = int32(_a_F_heap_abort_speculative_6)
						v82 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
						*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v82 + int32(1)
						v87 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[3]))
						v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+136))
						if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v87))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v91)) == int32(0) {
							v101 = base.B2i32(base.Ui32(v87) < base.Ui32(v91))
						} else {
							v101 = int32(base.Ui32(v87-v91) >> (uint(int32(31)) % 32))
						}
						v103 = v14 + int32(16)
						if v101 != 0 {
							v104 = v91
						} else {
							v104 = v87
						}
						v105 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
						if v105 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v50))) = v104
						} else {
							v108 = int32(3)
							if base.B2i32(base.Ui32(v104) < base.Ui32(v108))|base.B2i32(base.Ui32(v105) < base.Ui32(v108)) == int32(0) {
								if v104-v105 < int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v50))) = v104
								} else {
								}
							} else {
								if base.Ui32(v105) <= base.Ui32(v104) {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v50))) = v104
								}
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(0)
						v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+20)))
						v124 = v122 & int32(_a_F_heap_abort_speculative_7)
						*(*uint16)(unsafe.Add(mBase, uint32(v60)+20)) = uint16(v124)
						v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+18)))
						v128 = v126 & int32(_a_F_heap_abort_speculative_8)
						*(*uint16)(unsafe.Add(mBase, uint32(v60)+18)) = uint16(v128)
						v130 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
						*(*int32)(unsafe.Add(mBase, uint32(v60)+12)) = v130
						v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+4)))
						*(*uint16)(unsafe.Add(mBase, uint32(v60)+16)) = uint16(v132)
						F_MarkBufferDirty(m, v23)
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return
						} else {
							v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+118)))
							if v137 != int32(112) {
								v198 = int32(_a_F_heap_abort_speculative_6)
								v200 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
								*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v200 - int32(1)
								F_UnlockBuffer(m, v23)
								mBase = m.M
								v205 = m.ExcPending
								if v205 != 0 {
									return
								} else {
									v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+20)))
									if v206&int32(4) != 0 {
										F_heap_toast_delete(m, l0, v14+int32(12), int32(1))
										mBase = m.M
										v213 = m.ExcPending
										if v213 != 0 {
											return
										} else {
											F_ReleaseBuffer(m, v23)
											mBase = m.M
											v215 = m.ExcPending
											if v215 != 0 {
												return
											} else {
												F_pgstat_count_heap_delete(m, l0)
												mBase = m.M
												v217 = m.ExcPending
												if v217 != 0 {
													return
												} else {
													m.G0 = v14 + int32(32)
													return
												}
											}
										}
									} else {
										F_ReleaseBuffer(m, v23)
										mBase = m.M
										v215 = m.ExcPending
										if v215 != 0 {
											return
										} else {
											F_pgstat_count_heap_delete(m, l0)
											mBase = m.M
											v217 = m.ExcPending
											if v217 != 0 {
												return
											} else {
												m.G0 = v14 + int32(32)
												return
											}
										}
									}
								}
							} else {
								v141 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[4]))
								if v141 <= int32(0) {
									v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									if v144 != 0 {
										v198 = int32(_a_F_heap_abort_speculative_6)
										v200 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
										*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v200 - int32(1)
										F_UnlockBuffer(m, v23)
										mBase = m.M
										v205 = m.ExcPending
										if v205 != 0 {
											return
										} else {
											v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+20)))
											if v206&int32(4) != 0 {
												F_heap_toast_delete(m, l0, v14+int32(12), int32(1))
												mBase = m.M
												v213 = m.ExcPending
												if v213 != 0 {
													return
												} else {
													F_ReleaseBuffer(m, v23)
													mBase = m.M
													v215 = m.ExcPending
													if v215 != 0 {
														return
													} else {
														F_pgstat_count_heap_delete(m, l0)
														mBase = m.M
														v217 = m.ExcPending
														if v217 != 0 {
															return
														} else {
															m.G0 = v14 + int32(32)
															return
														}
													}
												}
											} else {
												F_ReleaseBuffer(m, v23)
												mBase = m.M
												v215 = m.ExcPending
												if v215 != 0 {
													return
												} else {
													F_pgstat_count_heap_delete(m, l0)
													mBase = m.M
													v217 = m.ExcPending
													if v217 != 0 {
														return
													} else {
														m.G0 = v14 + int32(32)
														return
													}
												}
											}
										}
									} else {
										v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										if v145 != 0 {
											v198 = int32(_a_F_heap_abort_speculative_6)
											v200 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
											*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v200 - int32(1)
											F_UnlockBuffer(m, v23)
											mBase = m.M
											v205 = m.ExcPending
											if v205 != 0 {
												return
											} else {
												v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+20)))
												if v206&int32(4) != 0 {
													F_heap_toast_delete(m, l0, v14+int32(12), int32(1))
													mBase = m.M
													v213 = m.ExcPending
													if v213 != 0 {
														return
													} else {
														F_ReleaseBuffer(m, v23)
														mBase = m.M
														v215 = m.ExcPending
														if v215 != 0 {
															return
														} else {
															F_pgstat_count_heap_delete(m, l0)
															mBase = m.M
															v217 = m.ExcPending
															if v217 != 0 {
																return
															} else {
																m.G0 = v14 + int32(32)
																return
															}
														}
													}
												} else {
													F_ReleaseBuffer(m, v23)
													mBase = m.M
													v215 = m.ExcPending
													if v215 != 0 {
														return
													} else {
														F_pgstat_count_heap_delete(m, l0)
														mBase = m.M
														v217 = m.ExcPending
														if v217 != 0 {
															return
														} else {
															m.G0 = v14 + int32(32)
															return
														}
													}
												}
											}
										} else {
											v146 = int32(8)
											*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)) = uint8(v146)
											v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+18)))
											v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+20)))
											v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+20)))
											*(*uint16)(unsafe.Add(mBase, uint32(v14)+8)) = uint16(v150)
											*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v16
											v157 = int32(1)
											v161 = int32(4)
											v176 = int32(base.Ui32(v148)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v149)>>(uint(v157)%32))&v146 | (int32(base.Ui32(v149)>>(uint(v161)%32))&v161 | (int32(base.Ui32(v149)>>(uint(int32(12))%32))&v157 | int32(base.Ui32(v149)>>(uint(int32(6))%32))&int32(2))))
											*(*uint8)(unsafe.Add(mBase, uint32(v14)+10)) = uint8(v176)
											F_XLogBeginInsert(m)
											mBase = m.M
											v179 = m.ExcPending
											if v179 != 0 {
												return
											} else {
												F_XLogRegisterData(m, v14+int32(4), int32(8))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return
												} else {
													F_XLogRegisterBuffer(m, int32(0), v23, int32(8))
													mBase = m.M
													v188 = m.ExcPending
													if v188 != 0 {
														return
													} else {
														v191 = F_XLogInsert(m, int32(10), int32(16))
														mBase = m.M
														v192 = m.ExcPending
														if v192 != 0 {
															return
														} else {
															*(*int64)(unsafe.Add(mBase, uint32(v42))) = base.I64_rotl(v191, int64(32))
															v198 = int32(_a_F_heap_abort_speculative_6)
															v200 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
															*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v200 - int32(1)
															F_UnlockBuffer(m, v23)
															mBase = m.M
															v205 = m.ExcPending
															if v205 != 0 {
																return
															} else {
																v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+20)))
																if v206&int32(4) != 0 {
																	F_heap_toast_delete(m, l0, v14+int32(12), int32(1))
																	mBase = m.M
																	v213 = m.ExcPending
																	if v213 != 0 {
																		return
																	} else {
																		F_ReleaseBuffer(m, v23)
																		mBase = m.M
																		v215 = m.ExcPending
																		if v215 != 0 {
																			return
																		} else {
																			F_pgstat_count_heap_delete(m, l0)
																			mBase = m.M
																			v217 = m.ExcPending
																			if v217 != 0 {
																				return
																			} else {
																				m.G0 = v14 + int32(32)
																				return
																			}
																		}
																	}
																} else {
																	F_ReleaseBuffer(m, v23)
																	mBase = m.M
																	v215 = m.ExcPending
																	if v215 != 0 {
																		return
																	} else {
																		F_pgstat_count_heap_delete(m, l0)
																		mBase = m.M
																		v217 = m.ExcPending
																		if v217 != 0 {
																			return
																		} else {
																			m.G0 = v14 + int32(32)
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
									v146 = int32(8)
									*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)) = uint8(v146)
									v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+18)))
									v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+20)))
									v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+20)))
									*(*uint16)(unsafe.Add(mBase, uint32(v14)+8)) = uint16(v150)
									*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v16
									v157 = int32(1)
									v161 = int32(4)
									v176 = int32(base.Ui32(v148)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v149)>>(uint(v157)%32))&v146 | (int32(base.Ui32(v149)>>(uint(v161)%32))&v161 | (int32(base.Ui32(v149)>>(uint(int32(12))%32))&v157 | int32(base.Ui32(v149)>>(uint(int32(6))%32))&int32(2))))
									*(*uint8)(unsafe.Add(mBase, uint32(v14)+10)) = uint8(v176)
									F_XLogBeginInsert(m)
									mBase = m.M
									v179 = m.ExcPending
									if v179 != 0 {
										return
									} else {
										F_XLogRegisterData(m, v14+int32(4), int32(8))
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
											return
										} else {
											F_XLogRegisterBuffer(m, int32(0), v23, int32(8))
											mBase = m.M
											v188 = m.ExcPending
											if v188 != 0 {
												return
											} else {
												v191 = F_XLogInsert(m, int32(10), int32(16))
												mBase = m.M
												v192 = m.ExcPending
												if v192 != 0 {
													return
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v42))) = base.I64_rotl(v191, int64(32))
													v198 = int32(_a_F_heap_abort_speculative_6)
													v200 = *(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2]))
													*(*int32)(unsafe.Add(mBase, _c_F_heap_abort_speculative[2])) = v200 - int32(1)
													F_UnlockBuffer(m, v23)
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return
													} else {
														v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+20)))
														if v206&int32(4) != 0 {
															F_heap_toast_delete(m, l0, v14+int32(12), int32(1))
															mBase = m.M
															v213 = m.ExcPending
															if v213 != 0 {
																return
															} else {
																F_ReleaseBuffer(m, v23)
																mBase = m.M
																v215 = m.ExcPending
																if v215 != 0 {
																	return
																} else {
																	F_pgstat_count_heap_delete(m, l0)
																	mBase = m.M
																	v217 = m.ExcPending
																	if v217 != 0 {
																		return
																	} else {
																		m.G0 = v14 + int32(32)
																		return
																	}
																}
															}
														} else {
															F_ReleaseBuffer(m, v23)
															mBase = m.M
															v215 = m.ExcPending
															if v215 != 0 {
																return
															} else {
																F_pgstat_count_heap_delete(m, l0)
																mBase = m.M
																v217 = m.ExcPending
																if v217 != 0 {
																	return
																} else {
																	m.G0 = v14 + int32(32)
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
					v224 = m.ExcPending
					if v224 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_heap_abort_speculative_9), int32(0))
						mBase = m.M
						v228 = m.ExcPending
						if v228 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_heap_abort_speculative_3), int32(_a_F_heap_abort_speculative_10), int32(_a_F_heap_abort_speculative_5))
							mBase = m.M
							v233 = m.ExcPending
							if v233 != 0 {
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	F_RelationIncrementReferenceCount(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = F_palloc(m, int32(700))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v19)+60)) = int64(0)
	if l1 == v21 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L55
	}
L5:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+28)))
	if v76&int32(5) != 0 {
		goto L26
	} else {
		goto L27
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = l5 & int32(-257)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v35 == int32(0) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = l5 & int32(-257)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v41 != int32(5) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_heap_beginscan[0]))
	if v45 <= int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_beginscan[1])))
	if v49&int32(1) == int32(0) {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+118)))
	if v55 != int32(112) {
		goto L4
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	if v45 <= int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v60 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L21
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v61 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	if base.Ui32(v62) < base.Ui32(int32(_a_F_heap_beginscan_0)) {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v65 == int32(0) {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+119)))
	switch v69 - int32(109) {
	case 0, 5:
		goto L24
	default:
		goto L4
	}
L24:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+112)))
	if v72 == int32(0) {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	goto L5
L26:
	;
	F_PredicateLockRelation(m, l0, l1)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v81
	if l4 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v84 = F_palloc(m, int32(16))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	v87 = int32(0)
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+100)) = v87
	if int32(0) < l2 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v87 = v84
	goto L32
L34:
	;
	v92 = F_palloc_mul(m, int32(56), l2)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	v95 = int32(0)
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v95
	F_initscan(m, v19, l3, int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	v95 = v92
	goto L36
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = int32(0)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	if v102&int32(17) != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = int32(0)
	m.G0 = v12 + int32(16)
	return v19
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v128
	v132 = int32(0)
	if base.B2i32(l5&int32(2048) == v132)|base.B2i32(v128 == v132) != 0 {
		goto L39
	} else {
		goto L50
	}
L41:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	if v111 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	if v102&int32(2) == int32(0) {
		goto L39
	} else {
		goto L48
	}
L44:
	;
	v112 = int32(140)
	goto L46
L45:
	;
	v112 = int32(141)
	goto L46
L46:
	;
	v114 = F_read_stream_begin_relation(m, int32(10), v106, v107, int32(0), v112, v19, int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v128 = v114
	goto L40
L48:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v126 = F_read_stream_begin_relation(m, int32(8), v121, v122, int32(0), int32(142), v19, int32(12))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v128 = v126
	goto L40
L50:
	;
	v138 = F_palloc0(m, int32(56))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v138
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+36)) = v138
	if v138 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+18)) = uint16(v143)
	goto L54
L53:
	;
	goto L54
L54:
	;
	goto L39
L55:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v161 + int32(4)
	F_errmsg(m, int32(_a_F_heap_beginscan_1), v12)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_heap_beginscan_2), int32(1223), int32(_a_F_heap_beginscan_3))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_compare_slots_2(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	v4 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+112))
	if v12 <= v4 {
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+120))
	v19 = int32(2)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17+base.I32_wrap_i64(l1)<<(uint(v19)%32))))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17+base.I32_wrap_i64(l0)<<(uint(v19)%32))))
	v35 = v4
	goto L6
L4:
	;
	return int32(-1)
L5:
	;
	v107 = int32(0)
	if v98 < v107 {
		goto L34
	} else {
		goto L35
	}
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+116))
	v42 = v39 + v35*int32(36)
	v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42)+10)))
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+6)))
	if v44 < v43 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	m.T0[v47].(func(*base.Module, int32, int32))(m, v27, v43)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v53 = v43 - int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v54))))
	v58 = v53 << (uint(int32(3)) % 32)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v58+v59)))
	v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+6)))
	if v62 < v43 {
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
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	m.T0[v65].(func(*base.Module, int32, int32))(m, v22, v43)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v53))))
	if v56&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v101 = v35 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l2)+112))
	if v101 < v102 {
		v35 = v101
		goto L6
	} else {
		goto L33
	}
L18:
	;
	if v70&int32(1) != 0 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v70&int32(1) != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+9)))
	if v75 == int32(0) {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	return int32(1)
L23:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+9)))
	if v82 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v85+v58)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v89 = m.T0[v88].(func(*base.Module, int64, int64, int32) int32)(m, v61, v87, v42)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L11
	} else {
		goto L27
	}
L26:
	;
	return int32(1)
L27:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+8)))
	if v91 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v89 < int32(0) {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	v98 = v89
	goto L30
L30:
	;
	if v98 != 0 {
		goto L5
	} else {
		goto L32
	}
L31:
	;
	v98 = int32(0) - v89
	goto L30
L32:
	;
	goto L17
L33:
	;
	goto L7
L34:
	;
	v111 = int32(1)
	goto L36
L35:
	;
	v111 = v107 - v98
	goto L36
L36:
	;
	return v111
}
func F_heap_compute_data_size(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	v4 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 <= v4 {
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
	v20 = int32(0)
	v24 = v4
	goto L4
L4:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+l2))))
	if v32 != 0 {
		v138 = v24
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return v138
L6:
	;
	v144 = v20 + int32(1)
	if v144 != v12 {
		v20 = v144
		v24 = v138
		goto L4
	} else {
		goto L37
	}
L7:
	;
	v34 = v20 << (uint(int32(3)) % 32)
	v36 = *(*int64)(unsafe.Add(mBase, uint32(l1+v34)))
	v37 = v34 + (l0 + int32(28))
	v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37)+2)))
	if v38 == int32(-1) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v138 = v131 + v132
	goto L6
L9:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v131 = int32(base.Ui32(v128) >> (uint(int32(2)) % 32))
	v132 = v95
	goto L8
L10:
	;
	v115 = int32(18)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v117 == v115 {
		goto L31
	} else {
		goto L32
	}
L11:
	;
	v41 = base.I32_wrap_i64(v36)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+6)))
	if v42&int32(1) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+5)))
	v105 = int32(0)
	v107 = (v24 + v101 - int32(1)) & (v105 - v101)
	if v105 < v38 {
		v131 = v38
		v132 = v107
		goto L8
	} else {
		goto L30
	}
L14:
	;
	if v62 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v62 = v61
	goto L14
L16:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v47&int32(3) != 0 {
		v62 = v47
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v52 = int32(base.Ui32(v50) >> (uint(int32(2)) % 32))
	if base.Ui32(int32(127)) < base.Ui32(v52-int32(3)) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v138 = v52 + v24 - int32(3)
	goto L6
L19:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v65&int32(254) != int32(2) {
		goto L10
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v85 = v62 & int32(1)
	if v85 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+5)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v36))+2))
	goto L23
L23:
	;
	v79 = F_EOH_get_flat_size(m, v78)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	v138 = (v24+v70-int32(1))&(int32(0)-v70) + v79
	goto L6
L26:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+5)))
	v95 = (v24 + v88 - int32(1)) & (int32(0) - v88)
	goto L28
L27:
	;
	v95 = v24
	goto L28
L28:
	;
	if v85 == int32(0) {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v131 = int32(base.Ui32(v62) >> (uint(int32(1)) % 32))
	v132 = v95
	goto L8
L30:
	;
	v111 = F_strlen(m, base.I32_wrap_i64(v36))
	mBase = m.M
	v131 = v111 + int32(1)
	v132 = v107
	goto L8
L31:
	;
	v120 = v115
	goto L33
L32:
	;
	v120 = int32(2)
	goto L33
L33:
	;
	if base.Ui32((v117-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v127 = int32(6)
	goto L36
L35:
	;
	v127 = v120
	goto L36
L36:
	;
	v131 = v127
	v132 = v24
	goto L8
L37:
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
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
	F_ReleaseBuffer(m, v6)
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v9 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L8
L10:
	;
	F_read_stream_end(m, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_RelationDecrementReferenceCount(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v15 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_pfree(m, v15)
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
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
	F_bms_free(m, v18)
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
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v21 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	F_pfree(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v24&int32(2) != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_UnregisterSnapshot(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v30 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	F_pfree(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	F_pfree(m, l0)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
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
					F_errfinish(m, int32(_a_F_heap_entry_is_visible_1), int32(1355), int32(_a_F_heap_entry_is_visible_2))
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
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	v15 = F_ReadBuffer(m, l0, v10|v11<<(uint(int32(16))%32))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		F_LockBufferInternal(m, v15, int32(1))
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
				F_UnlockReleaseBuffer(m, v15)
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return int32(0)
				} else {
					v114 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v114
					*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v114
					return v114
				}
			} else {
				v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+12)))
				if base.Ui32(int32(25)) <= base.Ui32(v43) {
					v51 = int32(base.Ui32(v43+int32(_a_F_heap_fetch_0)) >> (uint(int32(2)) % 32))
				} else {
					v51 = int32(0)
				}
				if base.Ui32(v51&int32(_a_F_heap_fetch_1)) < base.Ui32(v40) {
					F_UnlockReleaseBuffer(m, v15)
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return int32(0)
					} else {
						v114 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v114
						*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v114
						return v114
					}
				} else {
					v59 = v39 + v40<<(uint(int32(2))%32) + int32(20)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
					if v60&int32(_a_F_heap_fetch_2) != int32(_a_F_heap_fetch_3) {
						F_UnlockReleaseBuffer(m, v15)
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return int32(0)
						} else {
							v114 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v114
							*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v114
							return v114
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
										F_UnlockBuffer(m, v15)
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
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
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									F_UnlockBuffer(m, v15)
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										if l4 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v15
											return int32(0)
										} else {
											F_ReleaseBuffer(m, v15)
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return int32(0)
											} else {
												v114 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = v114
												*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v114
												return v114
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
	F_LockBufferInternal(m, v41, int32(1))
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
func F_heap_getsysattr(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
		return base.I64_extend_i32_u(v37)
	case 1, 3:
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
		v37 = v18
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v37)
	case 2:
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
		v37 = v16
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v37)
	case 4:
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v37 = v14
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v37)
	case 5:
		v37 = l0 + int32(4)
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v37)
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
			F_errmsg_internal(m, int32(_a_F_heap_getsysattr_0), v7)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_heap_getsysattr_1), int32(669), int32(_a_F_heap_getsysattr_2))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int64(0)
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
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v147 int32
	_ = v147
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v404 int32
	_ = v404
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v576 int32
	_ = v576
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v749 int32
	_ = v749
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v791 int64
	_ = v791
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v856 int32
	_ = v856
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v906 int32
	_ = v906
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1027 int32
	_ = v1027
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1041 int32
	_ = v1041
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1079 int32
	_ = v1079
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1187 int32
	_ = v1187
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1213 int32
	_ = v1213
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1239 int32
	_ = v1239
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1270 int32
	_ = v1270
	var v1277 int32
	_ = v1277
	var v1294 int32
	_ = v1294
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
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
		goto L14
	} else {
		goto L15
	}
L1:
	;
	if v76 != 0 {
		goto L321
	} else {
		goto L322
	}
L2:
	;
	v1294 = v1270
	v1301 = v1277
	v1303 = int32(0)
	goto L1
L3:
	;
	v1027 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v1027)
	if v1009&int32(4) != 0 {
		goto L266
	} else {
		goto L267
	}
L4:
	;
	v964 = int32(0)
	v966 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(v952 == v964)|base.B2i32(base.Ui32(v966) < base.Ui32(int32(3))) == v964 {
		goto L253
	} else {
		goto L254
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L33
	} else {
		goto L248
	}
L6:
	;
	if base.Ui32(v885) <= base.Ui32(v882) {
		v950 = v882
		v952 = v884
		goto L4
	} else {
		goto L247
	}
L7:
	;
	v419 = F_palloc_mul(m, int32(8), v190)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L33
	} else {
		goto L101
	}
L8:
	;
	v1294 = int32(0)
	v1301 = v404
	v1303 = v6
	goto L1
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L33
	} else {
		goto L97
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L33
	} else {
		goto L93
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L33
	} else {
		goto L89
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L33
	} else {
		goto L85
	}
L13:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	if v100&int32(_a_F_heap_prepare_freeze_tuple_0) != 0 {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = base.B2i32(base.Ui32(v42) < base.Ui32(int32(3)))
	if base.Ui32(v42) < base.Ui32(int32(3)) {
		v71 = v6
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v75 = int32(1)
	v76 = v6
	v77 = v37
	goto L16
L16:
	;
	if base.Ui32(v77&int32(_a_F_heap_prepare_freeze_tuple_1)) < base.Ui32(int32(_a_F_heap_prepare_freeze_tuple_2)) {
		v98 = v6
		goto L13
	} else {
		goto L22
	}
L17:
	;
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	v75 = v44
	v76 = v71
	v77 = v72
	goto L16
L18:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v45))&base.B2i32(v42-v45 < int32(0)) != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(v52) < base.Ui32(int32(3)))|base.B2i32(int32(0) <= v42-v52) != 0 {
		v71 = v6
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v59 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v59)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v62))&base.B2i32(v42-v62 <= int32(0)) != 0 {
		v71 = v59
		goto L17
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v42
	v71 = v59
	goto L17
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v82) < base.Ui32(int32(3)) {
		v98 = v6
		goto L13
	} else {
		goto L23
	}
L23:
	;
	v85 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v85)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v88))&base.B2i32(v82-v88 <= int32(0)) != 0 {
		v98 = v85
		goto L13
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v82
	v98 = v85
	goto L13
L25:
	;
	v103 = int32(2)
	if base.B2i32(v99 == int32(0))|base.B2i32(v100&int32(_a_F_heap_prepare_freeze_tuple_3) == int32(_a_F_heap_prepare_freeze_tuple_4)) != 0 {
		v1009 = v103
		v1013 = v6
		goto L3
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v99) {
		goto L72
	} else {
		goto L73
	}
L28:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v99-v111 < int32(0) {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	v116 = v100 & int32(128)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v99-v117 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v123 = F_MultiXactIdIsRunning(m, v99, base.B2i32(v116 != int32(0)))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v190 = F_GetMultiXactIdMembers(m, v99, v26+int32(168), base.B2i32(v116 != int32(0)))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L33
	} else {
		goto L50
	}
L33:
	;
	return int32(0)
L34:
	;
	if v123 != 0 {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	if v116 != 0 {
		v1009 = v103
		v1013 = v6
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v130 = F_GetMultiXactIdMembers(m, v99, v26+int32(172), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L33
	} else {
		goto L37
	}
L37:
	;
	if v130 <= int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v882 = v6
	v884 = v6
	v885 = v134
	goto L6
L39:
	;
	goto L40
L40:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v26)+172))
	v147 = v6
	goto L43
L41:
	;
	F_pfree(m, v135)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L33
	} else {
		goto L47
	}
L42:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v171 = v169
	goto L41
L43:
	;
	v161 = v135 + v147<<(uint(int32(3))%32)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v162) {
		goto L42
	} else {
		goto L45
	}
L44:
	;
	v171 = int32(0)
	goto L41
L45:
	;
	v166 = v147 + int32(1)
	if v166 != v130 {
		v147 = v166
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v176 = int32(3)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.B2i32(base.Ui32(v171) < base.Ui32(v176))|base.B2i32(base.Ui32(v178) < base.Ui32(v176)) != 0 {
		v882 = v171
		v884 = base.B2i32(base.Ui32(int32(2)) < base.Ui32(v171))
		v885 = v178
		goto L6
	} else {
		goto L48
	}
L48:
	;
	if v171-v178 < int32(0) {
		v906 = v171
		goto L5
	} else {
		goto L49
	}
L49:
	;
	v950 = v171
	v952 = int32(1)
	goto L4
L50:
	;
	if v190 <= int32(0) {
		v1009 = v103
		v1013 = v6
		goto L3
	} else {
		goto L51
	}
L51:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v26)+168))
	v208 = v195
	v209 = v6
	goto L52
L52:
	;
	v220 = int32(3)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v196+v209<<(uint(v220)%32))))
	if base.B2i32(base.Ui32(v194) < base.Ui32(v220))|base.B2i32(base.Ui32(v225) < base.Ui32(v220)) == int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v99-v249 < int32(0) {
		goto L7
	} else {
		goto L67
	}
L54:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v225))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v208)) != 0 {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	if int32(0) <= v225-v194 {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if base.Ui32(v225) < base.Ui32(v194) {
		goto L7
	} else {
		goto L59
	}
L58:
	;
	goto L7
L59:
	;
	goto L54
L60:
	;
	v244 = int32(base.Ui32(v225-v208) >> (uint(int32(31)) % 32))
	goto L62
L61:
	;
	v244 = base.B2i32(base.Ui32(v225) < base.Ui32(v208))
	goto L62
L62:
	;
	if v244 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v245 = v225
	goto L65
L64:
	;
	v245 = v208
	goto L65
L65:
	;
	v247 = v209 + int32(1)
	if v247 != v190 {
		v208 = v245
		v209 = v247
		goto L52
	} else {
		goto L66
	}
L66:
	;
	goto L53
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v245
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v99-v254 < int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v99
	goto L70
L69:
	;
	goto L70
L70:
	;
	F_pfree(m, v196)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L33
	} else {
		goto L71
	}
L71:
	;
	v404 = int32(0)
	goto L8
L72:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v264))&base.B2i32(v99-v264 < int32(0)) != 0 {
		goto L9
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if v99 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(v271) < base.Ui32(int32(3)))|base.B2i32(int32(0) <= v99-v271) != 0 {
		v404 = v6
		goto L8
	} else {
		goto L76
	}
L76:
	;
	v278 = int32(1)
	if v100&int32(128)|base.B2i32(v100&int32(80) == int32(64)) != 0 {
		v1270 = v278
		v1277 = v6
		goto L2
	} else {
		goto L77
	}
L77:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)))
	v288 = v286 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v288)
	v1270 = v278
	v1277 = v6
	goto L2
L78:
	;
	v1294 = int32(0)
	v1301 = int32(1)
	v1303 = v6
	goto L1
L79:
	;
	goto L80
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L33
	} else {
		goto L81
	}
L81:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L33
	} else {
		goto L82
	}
L82:
	;
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v301
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v99
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_5), v26+int32(16))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L33
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_6), int32(_a_F_heap_prepare_freeze_tuple_7), int32(_a_F_heap_prepare_freeze_tuple_8))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L33
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L33
	} else {
		goto L86
	}
L86:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+164)) = v321
	*(*int32)(unsafe.Add(mBase, uint32(v26)+160)) = v42
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_9), v26+int32(160))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L33
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_6), int32(_a_F_heap_prepare_freeze_tuple_10), int32(_a_F_heap_prepare_freeze_tuple_8))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L33
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L33
	} else {
		goto L90
	}
L90:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v341
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v99
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_11), v26+int32(32))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L33
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_6), int32(_a_F_heap_prepare_freeze_tuple_12), int32(_a_F_heap_prepare_freeze_tuple_13))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L33
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L33
	} else {
		goto L94
	}
L94:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v99
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_14), v26+int32(48))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L33
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_6), int32(_a_F_heap_prepare_freeze_tuple_15), int32(_a_F_heap_prepare_freeze_tuple_13))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L33
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L33
	} else {
		goto L98
	}
L98:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v381
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v99
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_16), v26)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L33
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_6), int32(_a_F_heap_prepare_freeze_tuple_17), int32(_a_F_heap_prepare_freeze_tuple_8))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L33
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	v421 = int32(0)
	v432 = v6
	v433 = v421
	v435 = v421
	v436 = v421
	v439 = v6
	goto L104
L102:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v26)+168))
	F_pfree(m, v849)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L33
	} else {
		goto L234
	}
L103:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L33
	} else {
		goto L230
	}
L104:
	;
	v447 = int32(3)
	v448 = v435 << (uint(v447) % 32)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v26)+168))
	v450 = v448 + v449
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v450)))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v450)+4))
	if base.Ui32(v452) <= base.Ui32(v447) {
		goto L109
	} else {
		goto L110
	}
L105:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L33
	} else {
		goto L225
	}
L106:
	;
	goto L105
L107:
	;
	v801 = v435 + int32(1)
	if v190 != v801 {
		v432 = v796
		v433 = v797
		v435 = v801
		v436 = v798
		v439 = v799
		goto L104
	} else {
		goto L224
	}
L108:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v26)+168))
	v791 = *(*int64)(unsafe.Add(mBase, uint32(v789+v448)))
	*(*int64)(unsafe.Add(mBase, uint32(v419+v436<<(uint(int32(3))%32)))) = v791
	v796 = v783
	v797 = v784
	v798 = v436 + int32(1)
	v799 = v785
	goto L107
L109:
	;
	if base.Ui32(v451) < base.Ui32(int32(3)) {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	goto L111
L111:
	;
	if v433 != 0 {
		goto L106
	} else {
		goto L167
	}
L112:
	;
	if v586 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L113:
	;
	v586 = int32(0)
	goto L112
L114:
	;
	goto L115
L115:
	;
	v466 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_freeze_tuple[0]))
	if v466 == v451 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v586 = int32(1)
	goto L112
L117:
	;
	goto L118
L118:
	;
	v470 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_freeze_tuple[1]))
	if v470 <= int32(0) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v586 = v576
	goto L112
L120:
	;
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_freeze_tuple[2]))
	if v474 == int32(0) {
		v576 = int32(0)
		goto L119
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v544 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_freeze_tuple[3]))
	v546 = int32(0)
	v549 = v470 - int32(1)
	goto L142
L123:
	;
	v479 = v474
	goto L124
L124:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v479)+20))
	if v485 == int32(4) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v576 = int32(0)
	goto L119
L126:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v479)+80))
	if v539 != 0 {
		v479 = v539
		goto L124
	} else {
		goto L141
	}
L127:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v479)))
	if v488 == int32(0) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v491 = int32(1)
	if v451 == v488 {
		v576 = v491
		goto L119
	} else {
		goto L129
	}
L129:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v479)+52))
	v495 = v493 - int32(1)
	if v495 < int32(0) {
		goto L126
	} else {
		goto L130
	}
L130:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v479)+48))
	v501 = int32(0)
	v504 = v495
	goto L131
L131:
	;
	v509 = int32(2)
	v510 = base.I32_div_s(v504-v501, v509)
	v511 = v510 + v501
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v498+v511<<(uint(v509)%32))))
	if v515 == v451 {
		v576 = v491
		goto L119
	} else {
		goto L133
	}
L132:
	;
	goto L126
L133:
	;
	v524 = base.B2i32(v515-v451 < int32(0)) | base.B2i32(base.Ui32(v515) < base.Ui32(int32(3)))
	if v524 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v525 = v511 + int32(1)
	goto L136
L135:
	;
	v525 = v501
	goto L136
L136:
	;
	if v524 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v528 = v504
	goto L139
L138:
	;
	v528 = v511 - int32(1)
	goto L139
L139:
	;
	if v525 <= v528 {
		v501 = v525
		v504 = v528
		goto L131
	} else {
		goto L140
	}
L140:
	;
	goto L132
L141:
	;
	goto L125
L142:
	;
	v554 = int32(2)
	v555 = base.I32_div_s(v549-v546, v554)
	v556 = v555 + v546
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v544+v556<<(uint(v554)%32))))
	v561 = base.B2i32(v560 == v451)
	if v560 == v451 {
		v576 = v561
		goto L119
	} else {
		goto L144
	}
L143:
	;
	v576 = v561
	goto L119
L144:
	;
	v564 = base.B2i32(base.Ui32(v560) < base.Ui32(v451))
	if base.Ui32(v560) < base.Ui32(v451) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v565 = v556 + int32(1)
	goto L147
L146:
	;
	v565 = v546
	goto L147
L147:
	;
	if base.Ui32(v560) < base.Ui32(v451) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v568 = v549
	goto L150
L149:
	;
	v568 = v556 - int32(1)
	goto L150
L150:
	;
	if v565 <= v568 {
		v546 = v565
		v549 = v568
		goto L142
	} else {
		goto L151
	}
L151:
	;
	goto L143
L152:
	;
	v589 = F_TransactionIdIsInProgress(m, v451)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L33
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v593 = int32(3)
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(v451) < base.Ui32(v593))|base.B2i32(base.Ui32(v595) < base.Ui32(v593)) == int32(0) {
		goto L158
	} else {
		goto L159
	}
L155:
	;
	if v589 == int32(0) {
		v796 = v432
		v797 = v433
		v798 = v436
		v799 = v439
		goto L107
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L33
	} else {
		goto L163
	}
L158:
	;
	if v451-v595 < int32(0) {
		goto L157
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	if base.Ui32(v451) < base.Ui32(v595) {
		goto L157
	} else {
		goto L162
	}
L161:
	;
	v783 = v432
	v784 = v433
	v785 = int32(1)
	goto L108
L162:
	;
	v783 = v432
	v784 = v433
	v785 = int32(1)
	goto L108
L163:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L33
	} else {
		goto L164
	}
L164:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+152)) = v614
	*(*int32)(unsafe.Add(mBase, uint32(v26)+148)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v26)+144)) = v99
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_18), v26+int32(144))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L33
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_6), int32(_a_F_heap_prepare_freeze_tuple_19), int32(_a_F_heap_prepare_freeze_tuple_13))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L33
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	if base.Ui32(v451) < base.Ui32(int32(3)) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v770 = int32(3)
	v772 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(v451) < base.Ui32(v770))|base.B2i32(base.Ui32(v772) < base.Ui32(v770)) == int32(0) {
		goto L219
	} else {
		goto L220
	}
L169:
	;
	if v759 != 0 {
		goto L209
	} else {
		goto L210
	}
L170:
	;
	v759 = int32(0)
	goto L169
L171:
	;
	goto L172
L172:
	;
	v639 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_freeze_tuple[0]))
	if v639 == v451 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v759 = int32(1)
	goto L169
L174:
	;
	goto L175
L175:
	;
	v643 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_freeze_tuple[1]))
	if v643 <= int32(0) {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v759 = v749
	goto L169
L177:
	;
	v647 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_freeze_tuple[2]))
	if v647 == int32(0) {
		v749 = int32(0)
		goto L176
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v717 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_freeze_tuple[3]))
	v719 = int32(0)
	v722 = v643 - int32(1)
	goto L199
L180:
	;
	v652 = v647
	goto L181
L181:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v652)+20))
	if v658 == int32(4) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v749 = int32(0)
	goto L176
L183:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v652)+80))
	if v712 != 0 {
		v652 = v712
		goto L181
	} else {
		goto L198
	}
L184:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v652)))
	if v661 == int32(0) {
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v664 = int32(1)
	if v451 == v661 {
		v749 = v664
		goto L176
	} else {
		goto L186
	}
L186:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v652)+52))
	v668 = v666 - int32(1)
	if v668 < int32(0) {
		goto L183
	} else {
		goto L187
	}
L187:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v652)+48))
	v674 = int32(0)
	v677 = v668
	goto L188
L188:
	;
	v682 = int32(2)
	v683 = base.I32_div_s(v677-v674, v682)
	v684 = v683 + v674
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v671+v684<<(uint(v682)%32))))
	if v688 == v451 {
		v749 = v664
		goto L176
	} else {
		goto L190
	}
L189:
	;
	goto L183
L190:
	;
	v697 = base.B2i32(v688-v451 < int32(0)) | base.B2i32(base.Ui32(v688) < base.Ui32(int32(3)))
	if v697 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v698 = v684 + int32(1)
	goto L193
L192:
	;
	v698 = v674
	goto L193
L193:
	;
	if v697 != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v701 = v677
	goto L196
L195:
	;
	v701 = v684 - int32(1)
	goto L196
L196:
	;
	if v698 <= v701 {
		v674 = v698
		v677 = v701
		goto L188
	} else {
		goto L197
	}
L197:
	;
	goto L189
L198:
	;
	goto L182
L199:
	;
	v727 = int32(2)
	v728 = base.I32_div_s(v722-v719, v727)
	v729 = v728 + v719
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v717+v729<<(uint(v727)%32))))
	v734 = base.B2i32(v733 == v451)
	if v733 == v451 {
		v749 = v734
		goto L176
	} else {
		goto L201
	}
L200:
	;
	v749 = v734
	goto L176
L201:
	;
	v737 = base.B2i32(base.Ui32(v733) < base.Ui32(v451))
	if base.Ui32(v733) < base.Ui32(v451) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v738 = v729 + int32(1)
	goto L204
L203:
	;
	v738 = v719
	goto L204
L204:
	;
	if base.Ui32(v733) < base.Ui32(v451) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v741 = v722
	goto L207
L206:
	;
	v741 = v729 - int32(1)
	goto L207
L207:
	;
	if v738 <= v741 {
		v719 = v738
		v722 = v741
		goto L199
	} else {
		goto L208
	}
L208:
	;
	goto L200
L209:
	;
	v768 = v432
	goto L168
L210:
	;
	goto L211
L211:
	;
	v760 = F_TransactionIdIsInProgress(m, v451)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L33
	} else {
		goto L212
	}
L212:
	;
	if v760 != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v768 = v432
	goto L168
L214:
	;
	goto L215
L215:
	;
	v764 = F_TransactionIdDidCommit(m, v451)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L33
	} else {
		goto L216
	}
L216:
	;
	if v764 == int32(0) {
		v796 = v432
		v797 = int32(0)
		v798 = v436
		v799 = v439
		goto L107
	} else {
		goto L217
	}
L217:
	;
	v768 = int32(1)
	goto L168
L218:
	;
	v783 = v768
	v784 = v451
	v785 = v439
	goto L108
L219:
	;
	if int32(0) <= v451-v772 {
		goto L218
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	if base.Ui32(v451) < base.Ui32(v772) {
		goto L103
	} else {
		goto L223
	}
L222:
	;
	goto L103
L223:
	;
	goto L218
L224:
	;
	goto L102
L225:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L33
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+128)) = v99
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_20), v26+int32(128))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L33
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+116)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v26)+112)) = v433
	F_errdetail_internal(m, int32(_a_F_heap_prepare_freeze_tuple_21), v26+int32(112))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L33
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_6), int32(_a_F_heap_prepare_freeze_tuple_22), int32(_a_F_heap_prepare_freeze_tuple_13))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L33
	} else {
		goto L229
	}
L229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L230:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L33
	} else {
		goto L231
	}
L231:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+104)) = v835
	*(*int32)(unsafe.Add(mBase, uint32(v26)+100)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v99
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_23), v26+int32(96))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L33
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_6), int32(_a_F_heap_prepare_freeze_tuple_24), int32(_a_F_heap_prepare_freeze_tuple_13))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L33
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	if v798 == int32(0) {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	F_pfree(m, v419)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L33
	} else {
		goto L246
	}
L236:
	;
	v869 = int32(2)
	v870 = int32(0)
	goto L235
L237:
	;
	goto L238
L238:
	;
	v856 = int32(0)
	if v799|base.B2i32(v797 == v856) == v856 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	if v796&int32(1) != 0 {
		goto L242
	} else {
		goto L243
	}
L240:
	;
	goto L241
L241:
	;
	v867 = F_MultiXactIdCreateFromMembers(m, v798, v419)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L33
	} else {
		goto L245
	}
L242:
	;
	v865 = int32(20)
	goto L244
L243:
	;
	v865 = int32(4)
	goto L244
L244:
	;
	v869 = v865
	v870 = v797
	goto L235
L245:
	;
	v869 = int32(8)
	v870 = v867
	goto L235
L246:
	;
	v1009 = v869
	v1013 = v870
	goto L3
L247:
	;
	v906 = v882
	goto L5
L248:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L33
	} else {
		goto L249
	}
L249:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = v927
	*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v906
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v99
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_25), v26-int32(-64))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L33
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_6), int32(_a_F_heap_prepare_freeze_tuple_26), int32(_a_F_heap_prepare_freeze_tuple_13))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L33
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L252:
	;
	v978 = F_TransactionIdDidCommit(m, v950)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L33
	} else {
		goto L258
	}
L253:
	;
	if v950-v966 < int32(0) {
		goto L252
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	if base.Ui32(v950) < base.Ui32(v966) {
		goto L252
	} else {
		goto L257
	}
L256:
	;
	v1009 = int32(4)
	v1013 = v950
	goto L3
L257:
	;
	v1009 = int32(4)
	v1013 = v950
	goto L3
L258:
	;
	if v978 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1009 = v103
	v1013 = int32(0)
	goto L3
L260:
	;
	goto L261
L261:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L33
	} else {
		goto L262
	}
L262:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L33
	} else {
		goto L263
	}
L263:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+88)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v26)+84)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = v99
	F_errmsg_internal(m, int32(_a_F_heap_prepare_freeze_tuple_23), v26+int32(80))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L33
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(_a_F_heap_prepare_freeze_tuple_6), int32(_a_F_heap_prepare_freeze_tuple_27), int32(_a_F_heap_prepare_freeze_tuple_13))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L33
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
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1013
	v1033 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
	v1035 = v1033 & int32(-7377)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v1035)
	if v1009&int32(16) != 0 {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	goto L268
L268:
	;
	if v1009&int32(8) == int32(0) {
		goto L272
	} else {
		goto L273
	}
L269:
	;
	v1041 = v1035 | int32(1024)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v1041)
	goto L271
L270:
	;
	goto L271
L271:
	;
	v1294 = int32(0)
	v1301 = int32(0)
	v1303 = v1027
	goto L1
L272:
	;
	v1270 = int32(1)
	v1277 = int32(0)
	goto L2
L273:
	;
	goto L274
L274:
	;
	v1050 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
	v1052 = v1050 & int32(_a_F_heap_prepare_freeze_tuple_28)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v1052)
	v1054 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	v1056 = v1054 & int32(_a_F_heap_prepare_freeze_tuple_29)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)) = uint16(v1056)
	v1058 = int32(0)
	v1062 = F_GetMultiXactIdMembers(m, v1013, v26+int32(172), v1058)
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L33
	} else {
		goto L276
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1013
	v1258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
	v1259 = v1258 | v1256
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v1259)
	v1261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	v1262 = v1261 | v1239
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)) = uint16(v1262)
	v1294 = int32(0)
	v1301 = v1058
	v1303 = v1027
	goto L1
L276:
	;
	if v1062 <= int32(0) {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1239 = int32(0)
	v1256 = int32(_a_F_heap_prepare_freeze_tuple_30)
	goto L275
L278:
	;
	goto L279
L279:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v26)+172))
	if v1062 == int32(1) {
		goto L282
	} else {
		goto L283
	}
L280:
	;
	F_pfree(m, v1068)
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L33
	} else {
		goto L309
	}
L281:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1068+v1160<<(uint(int32(3))%32))+4))
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1175<<(uint(int32(2))%32))+uint32(_c_F_heap_prepare_freeze_tuple[4])))
	if base.Ui32(v1154) < base.Ui32(v1178) {
		goto L303
	} else {
		goto L304
	}
L282:
	;
	v1071 = int32(0)
	v1154 = v1071
	v1155 = v1071
	v1157 = v1071
	v1160 = v1071
	goto L281
L283:
	;
	goto L284
L284:
	;
	v1079 = int32(0)
	v1089 = v1079
	v1090 = v1079
	v1092 = v1079
	v1095 = v1079
	v1097 = v1079
	goto L285
L285:
	;
	v1109 = v1068 + v1095<<(uint(int32(3))%32)
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1109)+4))
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1110<<(uint(int32(2))%32))+uint32(_c_F_heap_prepare_freeze_tuple[4])))
	if base.Ui32(v1089) < base.Ui32(v1113) {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	if v1062&int32(1) == int32(0) {
		v1194 = v1141
		v1195 = v1139
		v1197 = v1140
		goto L280
	} else {
		goto L302
	}
L287:
	;
	v1115 = v1113
	goto L289
L288:
	;
	v1115 = v1089
	goto L289
L289:
	;
	switch v1110 - int32(3) {
	case 0:
		goto L293
	case 1:
		v1122 = v1090
		goto L291
	case 2:
		goto L292
	default:
		v1124 = v1090
		v1125 = v1092
		goto L290
	}
L290:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1109)+12))
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1126<<(uint(int32(2))%32))+uint32(_c_F_heap_prepare_freeze_tuple[4])))
	switch v1126 - int32(3) {
	case 0:
		goto L297
	case 1:
		v1137 = v1124
		goto L295
	case 2:
		goto L296
	default:
		v1139 = v1124
		v1140 = v1125
		goto L294
	}
L291:
	;
	v1124 = v1122
	v1125 = int32(1)
	goto L290
L292:
	;
	v1122 = v1090 | int32(_a_F_heap_prepare_freeze_tuple_31)
	goto L291
L293:
	;
	v1124 = v1090 | int32(_a_F_heap_prepare_freeze_tuple_31)
	v1125 = v1092
	goto L290
L294:
	;
	if base.Ui32(v1115) < base.Ui32(v1129) {
		goto L298
	} else {
		goto L299
	}
L295:
	;
	v1139 = v1137
	v1140 = int32(1)
	goto L294
L296:
	;
	v1137 = v1124 | int32(_a_F_heap_prepare_freeze_tuple_31)
	goto L295
L297:
	;
	v1139 = v1124 | int32(_a_F_heap_prepare_freeze_tuple_31)
	v1140 = v1125
	goto L294
L298:
	;
	v1141 = v1129
	goto L300
L299:
	;
	v1141 = v1115
	goto L300
L300:
	;
	v1142 = int32(2)
	v1143 = v1095 + v1142
	v1145 = v1097 + v1142
	if v1145 != v1062&int32(2147483646) {
		v1089 = v1141
		v1090 = v1139
		v1092 = v1140
		v1095 = v1143
		v1097 = v1145
		goto L285
	} else {
		goto L301
	}
L301:
	;
	goto L286
L302:
	;
	v1154 = v1141
	v1155 = v1139
	v1157 = v1140
	v1160 = v1143
	goto L281
L303:
	;
	v1180 = v1178
	goto L305
L304:
	;
	v1180 = v1154
	goto L305
L305:
	;
	switch v1175 - int32(3) {
	case 0:
		goto L308
	case 1:
		v1187 = v1155
		goto L306
	case 2:
		goto L307
	default:
		v1194 = v1180
		v1195 = v1155
		v1197 = v1157
		goto L280
	}
L306:
	;
	v1194 = v1180
	v1195 = v1187
	v1197 = int32(1)
	goto L280
L307:
	;
	v1187 = v1155 | int32(_a_F_heap_prepare_freeze_tuple_31)
	goto L306
L308:
	;
	v1194 = v1180
	v1195 = v1155 | int32(_a_F_heap_prepare_freeze_tuple_31)
	v1197 = v1157
	goto L280
L309:
	;
	if v1194&int32(-2) == int32(2) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	if v1197&int32(1) != 0 {
		v1239 = v1195
		v1256 = int32(_a_F_heap_prepare_freeze_tuple_32)
		goto L275
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	if v1194 != 0 {
		goto L314
	} else {
		goto L315
	}
L313:
	;
	v1239 = v1195
	v1256 = int32(_a_F_heap_prepare_freeze_tuple_33)
	goto L275
L314:
	;
	v1225 = int32(_a_F_heap_prepare_freeze_tuple_0)
	goto L316
L315:
	;
	v1225 = int32(_a_F_heap_prepare_freeze_tuple_34)
	goto L316
L316:
	;
	if v1194 == int32(1) {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v1228 = int32(_a_F_heap_prepare_freeze_tuple_35)
	goto L319
L318:
	;
	v1228 = v1225
	goto L319
L319:
	;
	if v1197&int32(1) != 0 {
		v1239 = v1195
		v1256 = v1228
		goto L275
	} else {
		goto L320
	}
L320:
	;
	v1239 = v1195
	v1256 = v1228 | int32(128)
	goto L275
L321:
	;
	v1312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
	v1314 = v1312 | int32(768)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v1314)
	goto L323
L322:
	;
	goto L323
L323:
	;
	if v98 != 0 {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)))
	v1319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	if v1319&int32(_a_F_heap_prepare_freeze_tuple_2) != 0 {
		goto L327
	} else {
		goto L328
	}
L325:
	;
	goto L326
L326:
	;
	if v1294 != 0 {
		goto L330
	} else {
		goto L331
	}
L327:
	;
	v1322 = int32(4)
	goto L329
L328:
	;
	v1322 = int32(2)
	goto L329
L329:
	;
	v1323 = v1316 | v1322
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v1323)
	goto L326
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	v1327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	v1329 = v1327 & int32(_a_F_heap_prepare_freeze_tuple_36)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)) = uint16(v1329)
	v1331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
	v1335 = v1331&int32(_a_F_heap_prepare_freeze_tuple_28) | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v1335)
	goto L332
L331:
	;
	goto L332
L332:
	;
	v1339 = (v75 | v76) & (v1294 | v1301)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v1339)
	if v1301&v75 != 0 {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	m.G0 = v26 + int32(176)
	return v76 | (v1303 | v98) | v1294
L334:
	;
	v1342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v1342&int32(1) != 0 {
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v1349 = F_heap_tuple_should_freeze(m, l0, l1, l2+int32(16), l2+int32(20))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L33
	} else {
		goto L336
	}
L336:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v1349)
	goto L333
}
func F_heap_prepare_pagescan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v458 int32
	_ = v458
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v501 int32
	_ = v501
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(_a_F_heap_prepare_pagescan_0)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_heap_page_prune_opt(m, v23, v24, l0+int32(104), int32(base.Ui32(v27&int32(1024))>>(uint(int32(10))%32)))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_LockBufferInternal(m, v24, int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v24 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v55) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_pagescan[0]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40+(v24^int32(-1))<<(uint(int32(2))%32))))
	v54 = v46
	goto L4
L6:
	;
	goto L7
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_pagescan[1]))
	v54 = v48 + v24<<(uint(int32(13))%32) + int32(-8192)
	goto L4
L8:
	;
	v65 = int32(base.Ui32(v55+int32(_a_F_heap_prepare_pagescan_1))>>(uint(int32(2))%32)) & int32(_a_F_heap_prepare_pagescan_2)
	goto L10
L9:
	;
	v65 = int32(0)
	goto L10
L10:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+10)))
	if v66&int32(4) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v484
	F_UnlockBuffer(m, v24)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L90
	}
L12:
	;
	if v65 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L13:
	;
	if v267 <= int32(0) {
		v484 = v267
		goto L11
	} else {
		goto L69
	}
L14:
	;
	if v65 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L15:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v72 = F_CheckForSerializableConflictOutNeeded(m, v71, v21)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+29)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v78 = F_CheckForSerializableConflictOutNeeded(m, v77, v21)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	if v72 == int32(0) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L12
L20:
	;
	if v76 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v78 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	if v78 != 0 {
		goto L12
	} else {
		goto L57
	}
L24:
	;
	if v65 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	if v65 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L27:
	;
	v484 = int32(0)
	goto L11
L28:
	;
	goto L29
L29:
	;
	v87 = int32(1)
	v89 = l0 + int32(116)
	v91 = v54 + int32(20)
	v95 = (v65 + v87) & int32(_a_F_heap_prepare_pagescan_2)
	if base.Ui32(v95) < base.Ui32(int32(3)) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v91+v177<<(uint(int32(2))%32))))
	if v193&int32(_a_F_heap_prepare_pagescan_3) != int32(_a_F_heap_prepare_pagescan_4) {
		v484 = v175
		goto L11
	} else {
		goto L47
	}
L31:
	;
	v175 = int32(0)
	v176 = int32(1)
	v177 = v87
	goto L30
L32:
	;
	goto L33
L33:
	;
	v100 = int32(2)
	if base.Ui32(v95) <= base.Ui32(v100) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v103 = v100
	goto L36
L35:
	;
	v103 = v95
	goto L36
L36:
	;
	v104 = int32(1)
	v105 = v103 - v104
	v110 = int32(0)
	v114 = v110
	v115 = v104
	v116 = v87
	v120 = v110
	goto L37
L37:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v91+v116<<(uint(int32(2))%32))))
	if v132&int32(_a_F_heap_prepare_pagescan_3) == int32(_a_F_heap_prepare_pagescan_4) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v105&v104 == int32(0) {
		v484 = v162
		goto L11
	} else {
		goto L46
	}
L39:
	;
	v137 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v89+v114<<(uint(v137)%32)))) = uint16(v115)
	v143 = v114 + v137
	goto L41
L40:
	;
	v143 = v114
	goto L41
L41:
	;
	v145 = v115 + int32(1)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v91+v145&int32(_a_F_heap_prepare_pagescan_2)<<(uint(int32(2))%32))))
	if v151&int32(_a_F_heap_prepare_pagescan_3) == int32(_a_F_heap_prepare_pagescan_4) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v156 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v89+v143<<(uint(v156)%32)))) = uint16(v145)
	v162 = v143 + v156
	goto L44
L43:
	;
	v162 = v143
	goto L44
L44:
	;
	v163 = int32(2)
	v164 = v115 + v163
	v165 = int32(_a_F_heap_prepare_pagescan_2)
	v166 = v164 & v165
	v168 = v120 + v163
	if v168&v165 != v105&int32(_a_F_heap_prepare_pagescan_5) {
		v114 = v162
		v115 = v164
		v116 = v166
		v120 = v168
		goto L37
	} else {
		goto L45
	}
L45:
	;
	goto L38
L46:
	;
	v175 = v162
	v176 = v164
	v177 = v166
	goto L30
L47:
	;
	v198 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v89+v175<<(uint(v198)%32)))) = uint16(v176)
	v484 = v175 + v198
	goto L11
L48:
	;
	v484 = int32(0)
	goto L11
L49:
	;
	goto L50
L50:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+56))
	v212 = int32(base.Ui32(v22) >> (uint(int32(16)) % 32))
	v218 = int32(1)
	v221 = int32(0)
	v222 = v218
	v223 = v218
	goto L51
L51:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v54+int32(20)+v223<<(uint(int32(2))%32))))
	if v239&int32(_a_F_heap_prepare_pagescan_3) == int32(_a_F_heap_prepare_pagescan_4) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L13
L53:
	;
	v246 = v19 + v221*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v246)+12)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v246))) = int32(base.Ui32(v239) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+16)) = v54 + v239&int32(_a_F_heap_prepare_pagescan_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v246)+8)) = uint16(v222)
	*(*uint16)(unsafe.Add(mBase, uint32(v246)+6)) = uint16(v22)
	*(*uint16)(unsafe.Add(mBase, uint32(v246)+4)) = uint16(v212)
	v259 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v221+(v19+int32(_a_F_heap_prepare_pagescan_7))))) = uint8(v259)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(116)+v221<<(uint(v259)%32)))) = uint16(v222)
	v267 = v221 + v259
	goto L55
L54:
	;
	v267 = v221
	goto L55
L55:
	;
	v270 = v222 + int32(1)
	v272 = v270 & int32(_a_F_heap_prepare_pagescan_2)
	if base.Ui32(v272) <= base.Ui32(v65) {
		v221 = v267
		v222 = v270
		v223 = v272
		goto L51
	} else {
		goto L56
	}
L56:
	;
	goto L52
L57:
	;
	goto L14
L58:
	;
	v350 = F_HeapTupleSatisfiesMVCCBatch(m, v21, v24, v335, v19, l0+int32(116))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L68
	}
L59:
	;
	v335 = v2
	goto L58
L60:
	;
	goto L61
L61:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+56))
	v280 = int32(base.Ui32(v22) >> (uint(int32(16)) % 32))
	v283 = int32(1)
	v286 = v283
	v287 = v283
	v288 = v2
	goto L62
L62:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v54+int32(20)+v286<<(uint(int32(2))%32))))
	if v304&int32(_a_F_heap_prepare_pagescan_3) == int32(_a_F_heap_prepare_pagescan_4) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v335 = v326
	goto L58
L64:
	;
	v311 = v19 + v288*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v311)+12)) = v278
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = int32(base.Ui32(v304) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v311)+16)) = v54 + v304&int32(_a_F_heap_prepare_pagescan_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v311)+8)) = uint16(v287)
	*(*uint16)(unsafe.Add(mBase, uint32(v311)+6)) = uint16(v22)
	*(*uint16)(unsafe.Add(mBase, uint32(v311)+4)) = uint16(v280)
	v326 = v288 + int32(1)
	goto L66
L65:
	;
	v326 = v288
	goto L66
L66:
	;
	v328 = v287 + int32(1)
	v330 = v328 & int32(_a_F_heap_prepare_pagescan_2)
	if base.Ui32(v330) <= base.Ui32(v65) {
		v286 = v330
		v287 = v328
		v288 = v326
		goto L62
	} else {
		goto L67
	}
L67:
	;
	goto L63
L68:
	;
	v484 = v350
	goto L11
L69:
	;
	v359 = int32(0)
	goto L70
L70:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359+(v19+int32(_a_F_heap_prepare_pagescan_7))))))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_HeapCheckForSerializableConflictOut(m, v374, v375, v19+v359*int32(20), v24, v21)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L72
	}
L71:
	;
	v484 = v267
	goto L11
L72:
	;
	v382 = v359 + int32(1)
	if v382 != v267 {
		v359 = v382
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v390 = F_HeapTupleSatisfiesMVCCBatch(m, v21, v24, int32(0), v19, l0+int32(116))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+56))
	v395 = int32(base.Ui32(v22) >> (uint(int32(16)) % 32))
	v398 = int32(1)
	v401 = v398
	v402 = v398
	v403 = v2
	goto L78
L77:
	;
	v484 = v390
	goto L11
L78:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v54+int32(20)+v401<<(uint(int32(2))%32))))
	if v419&int32(_a_F_heap_prepare_pagescan_3) == int32(_a_F_heap_prepare_pagescan_4) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v449 = F_HeapTupleSatisfiesMVCCBatch(m, v21, v24, v441, v19, l0+int32(116))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L84
	}
L80:
	;
	v426 = v19 + v403*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v426)+12)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v426))) = int32(base.Ui32(v419) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v426)+16)) = v54 + v419&int32(_a_F_heap_prepare_pagescan_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v426)+8)) = uint16(v402)
	*(*uint16)(unsafe.Add(mBase, uint32(v426)+6)) = uint16(v22)
	*(*uint16)(unsafe.Add(mBase, uint32(v426)+4)) = uint16(v395)
	v441 = v403 + int32(1)
	goto L82
L81:
	;
	v441 = v403
	goto L82
L82:
	;
	v443 = v402 + int32(1)
	v445 = v443 & int32(_a_F_heap_prepare_pagescan_2)
	if base.Ui32(v445) <= base.Ui32(v65) {
		v401 = v445
		v402 = v443
		v403 = v441
		goto L78
	} else {
		goto L83
	}
L83:
	;
	goto L79
L84:
	;
	if v441 <= int32(0) {
		v484 = v449
		goto L11
	} else {
		goto L85
	}
L85:
	;
	v458 = int32(0)
	goto L86
L86:
	;
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458+(v19+int32(_a_F_heap_prepare_pagescan_7))))))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_HeapCheckForSerializableConflictOut(m, v473, v474, v19+v458*int32(20), v24, v21)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L88
	}
L87:
	;
	v484 = v449
	goto L11
L88:
	;
	v481 = v458 + int32(1)
	if v481 != v441 {
		v458 = v481
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	m.G0 = v19 + int32(_a_F_heap_prepare_pagescan_0)
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
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
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
				if v26 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v19 & int32(-257)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v19 | int32(256)
				}
			}
		}
	}
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v36 != 0 {
		F_ReleaseBuffer(m, v36)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(0)
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			if v41 != 0 {
				F_ReleaseBuffer(m, v41)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					if v46 != 0 {
						F_read_stream_reset(m, v46)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							F_initscan(m, l0, l1, int32(1))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						F_initscan(m, l0, l1, int32(1))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
				if v46 != 0 {
					F_read_stream_reset(m, v46)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_initscan(m, l0, l1, int32(1))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					F_initscan(m, l0, l1, int32(1))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		if v41 != 0 {
			F_ReleaseBuffer(m, v41)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
				if v46 != 0 {
					F_read_stream_reset(m, v46)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_initscan(m, l0, l1, int32(1))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					F_initscan(m, l0, l1, int32(1))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
			if v46 != 0 {
				F_read_stream_reset(m, v46)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					F_initscan(m, l0, l1, int32(1))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				F_initscan(m, l0, l1, int32(1))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
