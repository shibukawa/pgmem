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
	v19 = *(*int32)(unsafe.Add(mBase, _consts[38]))
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
	v23 = *(*int32)(unsafe.Add(mBase, _consts[39]))
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
	v27 = *(*int32)(unsafe.Add(mBase, _consts[40]))
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
	v89 = *(*int32)(unsafe.Add(mBase, _consts[41]))
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
	v136 = *(*int32)(unsafe.Add(mBase, _consts[1290]))
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
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
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
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
		v21 = F_ReadBuffer(m, l0, v16|v17<<(uint(int32(16))%32))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			if v21 < int32(0) {
				v26 = *(*int32)(unsafe.Add(mBase, _consts[12]))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v26+(v21^int32(-1))<<(uint(int32(2))%32))))
				v40 = v32
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, _consts[13]))
				v40 = v34 + v21<<(uint(int32(13))%32) + int32(-8192)
			}
			F_LockBuffer(m, v21, int32(2))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				v45 = l1 + int32(4)
				v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45))))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v47
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v46<<(uint(int32(2))%32)+v40)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v52) >> (uint(int32(17)) % 32))
				v58 = v40 + v52&int32(32767)
				*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v58
				v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45))))
				*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)) = uint16(v60)
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v62
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
				if v14 == v64 {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+68))
					if v69 != int32(99) {
						v72 = F_isTempToastNamespace(m, v69)
						mBase = m.M
						v73 = v72
					} else {
						v73 = int32(1)
					}
					if v73 == int32(0) {
						v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+16)))
						if v76 != int32(65534) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v240 = m.ExcPending
							if v240 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(403195), int32(0))
								mBase = m.M
								v244 = m.ExcPending
								if v244 != 0 {
									return
								} else {
									F_errfinish(m, int32(522199), int32(6224), int32(360639))
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
							v80 = v12 + int32(16)
							v81 = int32(4556740)
							v83 = *(*int32)(unsafe.Add(mBase, _consts[14]))
							*(*int32)(unsafe.Add(mBase, _consts[14])) = v83 + int32(1)
							v88 = *(*int32)(unsafe.Add(mBase, _consts[43]))
							v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+136))
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v90))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v88)) == int32(0) {
								v102 = base.B2i32(base.Ui32(v88) < base.Ui32(v90))
							} else {
								v102 = int32(base.Ui32(v88-v90) >> (uint(int32(31)) % 32))
							}
							v104 = *(*int32)(unsafe.Add(mBase, _consts[43]))
							if v102 != 0 {
								v105 = v90
							} else {
								v105 = v104
							}
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
							if v106 != 0 {
								if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v106))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v105)) == int32(0) {
									v118 = base.B2i32(base.Ui32(v105) < base.Ui32(v106))
								} else {
									v118 = int32(base.Ui32(v105-v106) >> (uint(int32(31)) % 32))
								}
								if v118 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v105
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v105
							}
							*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(0)
							v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+20)))
							v126 = v124 & int32(9007)
							*(*uint16)(unsafe.Add(mBase, uint32(v58)+20)) = uint16(v126)
							v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+18)))
							v130 = v128 & int32(57343)
							*(*uint16)(unsafe.Add(mBase, uint32(v58)+18)) = uint16(v130)
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
							*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v132
							v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+4)))
							*(*uint16)(unsafe.Add(mBase, uint32(v58)+16)) = uint16(v134)
							F_MarkBufferDirty(m, v21)
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return
							} else {
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+118)))
								if v139 != int32(112) {
									v200 = int32(4556740)
									v202 = *(*int32)(unsafe.Add(mBase, _consts[14]))
									*(*int32)(unsafe.Add(mBase, _consts[14])) = v202 - int32(1)
									F_LockBuffer(m, v21, int32(0))
									mBase = m.M
									v208 = m.ExcPending
									if v208 != 0 {
										return
									} else {
										v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+20)))
										if v209&int32(4) != 0 {
											F_heap_toast_delete(m, l0, v12+int32(12), int32(1))
											mBase = m.M
											v216 = m.ExcPending
											if v216 != 0 {
												return
											} else {
												F_ReleaseBuffer(m, v21)
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
														m.G0 = v12 + int32(32)
														return
													}
												}
											}
										} else {
											F_ReleaseBuffer(m, v21)
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
													m.G0 = v12 + int32(32)
													return
												}
											}
										}
									}
								} else {
									v143 = *(*int32)(unsafe.Add(mBase, _consts[15]))
									if v143 <= int32(0) {
										v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										if v146 != 0 {
											v200 = int32(4556740)
											v202 = *(*int32)(unsafe.Add(mBase, _consts[14]))
											*(*int32)(unsafe.Add(mBase, _consts[14])) = v202 - int32(1)
											F_LockBuffer(m, v21, int32(0))
											mBase = m.M
											v208 = m.ExcPending
											if v208 != 0 {
												return
											} else {
												v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+20)))
												if v209&int32(4) != 0 {
													F_heap_toast_delete(m, l0, v12+int32(12), int32(1))
													mBase = m.M
													v216 = m.ExcPending
													if v216 != 0 {
														return
													} else {
														F_ReleaseBuffer(m, v21)
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
																m.G0 = v12 + int32(32)
																return
															}
														}
													}
												} else {
													F_ReleaseBuffer(m, v21)
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
															m.G0 = v12 + int32(32)
															return
														}
													}
												}
											}
										} else {
											v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											if v147 != 0 {
												v200 = int32(4556740)
												v202 = *(*int32)(unsafe.Add(mBase, _consts[14]))
												*(*int32)(unsafe.Add(mBase, _consts[14])) = v202 - int32(1)
												F_LockBuffer(m, v21, int32(0))
												mBase = m.M
												v208 = m.ExcPending
												if v208 != 0 {
													return
												} else {
													v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+20)))
													if v209&int32(4) != 0 {
														F_heap_toast_delete(m, l0, v12+int32(12), int32(1))
														mBase = m.M
														v216 = m.ExcPending
														if v216 != 0 {
															return
														} else {
															F_ReleaseBuffer(m, v21)
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
																	m.G0 = v12 + int32(32)
																	return
																}
															}
														}
													} else {
														F_ReleaseBuffer(m, v21)
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
																m.G0 = v12 + int32(32)
																return
															}
														}
													}
												}
											} else {
												v148 = int32(8)
												*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)) = uint8(v148)
												v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+18)))
												v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+20)))
												v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)))
												*(*uint16)(unsafe.Add(mBase, uint32(v12)+8)) = uint16(v152)
												*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v14
												v159 = int32(1)
												v163 = int32(4)
												v178 = int32(base.Ui32(v150)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v151)>>(uint(v159)%32))&v148 | (int32(base.Ui32(v151)>>(uint(v163)%32))&v163 | (int32(base.Ui32(v151)>>(uint(int32(12))%32))&v159 | int32(base.Ui32(v151)>>(uint(int32(6))%32))&int32(2))))
												*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)) = uint8(v178)
												F_XLogBeginInsert(m)
												mBase = m.M
												v181 = m.ExcPending
												if v181 != 0 {
													return
												} else {
													F_XLogRegisterData(m, v12+int32(4), int32(8))
													mBase = m.M
													v186 = m.ExcPending
													if v186 != 0 {
														return
													} else {
														F_XLogRegisterBuffer(m, int32(0), v21, int32(8))
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
																*(*int64)(unsafe.Add(mBase, uint32(v40))) = base.I64_rotr(v193, int64(32))
																v200 = int32(4556740)
																v202 = *(*int32)(unsafe.Add(mBase, _consts[14]))
																*(*int32)(unsafe.Add(mBase, _consts[14])) = v202 - int32(1)
																F_LockBuffer(m, v21, int32(0))
																mBase = m.M
																v208 = m.ExcPending
																if v208 != 0 {
																	return
																} else {
																	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+20)))
																	if v209&int32(4) != 0 {
																		F_heap_toast_delete(m, l0, v12+int32(12), int32(1))
																		mBase = m.M
																		v216 = m.ExcPending
																		if v216 != 0 {
																			return
																		} else {
																			F_ReleaseBuffer(m, v21)
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
																					m.G0 = v12 + int32(32)
																					return
																				}
																			}
																		}
																	} else {
																		F_ReleaseBuffer(m, v21)
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
																				m.G0 = v12 + int32(32)
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
										*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)) = uint8(v148)
										v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+18)))
										v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+20)))
										v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)))
										*(*uint16)(unsafe.Add(mBase, uint32(v12)+8)) = uint16(v152)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v14
										v159 = int32(1)
										v163 = int32(4)
										v178 = int32(base.Ui32(v150)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v151)>>(uint(v159)%32))&v148 | (int32(base.Ui32(v151)>>(uint(v163)%32))&v163 | (int32(base.Ui32(v151)>>(uint(int32(12))%32))&v159 | int32(base.Ui32(v151)>>(uint(int32(6))%32))&int32(2))))
										*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)) = uint8(v178)
										F_XLogBeginInsert(m)
										mBase = m.M
										v181 = m.ExcPending
										if v181 != 0 {
											return
										} else {
											F_XLogRegisterData(m, v12+int32(4), int32(8))
											mBase = m.M
											v186 = m.ExcPending
											if v186 != 0 {
												return
											} else {
												F_XLogRegisterBuffer(m, int32(0), v21, int32(8))
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
														*(*int64)(unsafe.Add(mBase, uint32(v40))) = base.I64_rotr(v193, int64(32))
														v200 = int32(4556740)
														v202 = *(*int32)(unsafe.Add(mBase, _consts[14]))
														*(*int32)(unsafe.Add(mBase, _consts[14])) = v202 - int32(1)
														F_LockBuffer(m, v21, int32(0))
														mBase = m.M
														v208 = m.ExcPending
														if v208 != 0 {
															return
														} else {
															v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+20)))
															if v209&int32(4) != 0 {
																F_heap_toast_delete(m, l0, v12+int32(12), int32(1))
																mBase = m.M
																v216 = m.ExcPending
																if v216 != 0 {
																	return
																} else {
																	F_ReleaseBuffer(m, v21)
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
																			m.G0 = v12 + int32(32)
																			return
																		}
																	}
																}
															} else {
																F_ReleaseBuffer(m, v21)
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
																		m.G0 = v12 + int32(32)
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
						v80 = v12 + int32(16)
						v81 = int32(4556740)
						v83 = *(*int32)(unsafe.Add(mBase, _consts[14]))
						*(*int32)(unsafe.Add(mBase, _consts[14])) = v83 + int32(1)
						v88 = *(*int32)(unsafe.Add(mBase, _consts[43]))
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+136))
						if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v90))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v88)) == int32(0) {
							v102 = base.B2i32(base.Ui32(v88) < base.Ui32(v90))
						} else {
							v102 = int32(base.Ui32(v88-v90) >> (uint(int32(31)) % 32))
						}
						v104 = *(*int32)(unsafe.Add(mBase, _consts[43]))
						if v102 != 0 {
							v105 = v90
						} else {
							v105 = v104
						}
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
						if v106 != 0 {
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v106))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v105)) == int32(0) {
								v118 = base.B2i32(base.Ui32(v105) < base.Ui32(v106))
							} else {
								v118 = int32(base.Ui32(v105-v106) >> (uint(int32(31)) % 32))
							}
							if v118 == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v105
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v105
						}
						*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(0)
						v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+20)))
						v126 = v124 & int32(9007)
						*(*uint16)(unsafe.Add(mBase, uint32(v58)+20)) = uint16(v126)
						v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+18)))
						v130 = v128 & int32(57343)
						*(*uint16)(unsafe.Add(mBase, uint32(v58)+18)) = uint16(v130)
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
						*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v132
						v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+4)))
						*(*uint16)(unsafe.Add(mBase, uint32(v58)+16)) = uint16(v134)
						F_MarkBufferDirty(m, v21)
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return
						} else {
							v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+118)))
							if v139 != int32(112) {
								v200 = int32(4556740)
								v202 = *(*int32)(unsafe.Add(mBase, _consts[14]))
								*(*int32)(unsafe.Add(mBase, _consts[14])) = v202 - int32(1)
								F_LockBuffer(m, v21, int32(0))
								mBase = m.M
								v208 = m.ExcPending
								if v208 != 0 {
									return
								} else {
									v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+20)))
									if v209&int32(4) != 0 {
										F_heap_toast_delete(m, l0, v12+int32(12), int32(1))
										mBase = m.M
										v216 = m.ExcPending
										if v216 != 0 {
											return
										} else {
											F_ReleaseBuffer(m, v21)
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
													m.G0 = v12 + int32(32)
													return
												}
											}
										}
									} else {
										F_ReleaseBuffer(m, v21)
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
												m.G0 = v12 + int32(32)
												return
											}
										}
									}
								}
							} else {
								v143 = *(*int32)(unsafe.Add(mBase, _consts[15]))
								if v143 <= int32(0) {
									v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									if v146 != 0 {
										v200 = int32(4556740)
										v202 = *(*int32)(unsafe.Add(mBase, _consts[14]))
										*(*int32)(unsafe.Add(mBase, _consts[14])) = v202 - int32(1)
										F_LockBuffer(m, v21, int32(0))
										mBase = m.M
										v208 = m.ExcPending
										if v208 != 0 {
											return
										} else {
											v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+20)))
											if v209&int32(4) != 0 {
												F_heap_toast_delete(m, l0, v12+int32(12), int32(1))
												mBase = m.M
												v216 = m.ExcPending
												if v216 != 0 {
													return
												} else {
													F_ReleaseBuffer(m, v21)
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
															m.G0 = v12 + int32(32)
															return
														}
													}
												}
											} else {
												F_ReleaseBuffer(m, v21)
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
														m.G0 = v12 + int32(32)
														return
													}
												}
											}
										}
									} else {
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										if v147 != 0 {
											v200 = int32(4556740)
											v202 = *(*int32)(unsafe.Add(mBase, _consts[14]))
											*(*int32)(unsafe.Add(mBase, _consts[14])) = v202 - int32(1)
											F_LockBuffer(m, v21, int32(0))
											mBase = m.M
											v208 = m.ExcPending
											if v208 != 0 {
												return
											} else {
												v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+20)))
												if v209&int32(4) != 0 {
													F_heap_toast_delete(m, l0, v12+int32(12), int32(1))
													mBase = m.M
													v216 = m.ExcPending
													if v216 != 0 {
														return
													} else {
														F_ReleaseBuffer(m, v21)
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
																m.G0 = v12 + int32(32)
																return
															}
														}
													}
												} else {
													F_ReleaseBuffer(m, v21)
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
															m.G0 = v12 + int32(32)
															return
														}
													}
												}
											}
										} else {
											v148 = int32(8)
											*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)) = uint8(v148)
											v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+18)))
											v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+20)))
											v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)))
											*(*uint16)(unsafe.Add(mBase, uint32(v12)+8)) = uint16(v152)
											*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v14
											v159 = int32(1)
											v163 = int32(4)
											v178 = int32(base.Ui32(v150)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v151)>>(uint(v159)%32))&v148 | (int32(base.Ui32(v151)>>(uint(v163)%32))&v163 | (int32(base.Ui32(v151)>>(uint(int32(12))%32))&v159 | int32(base.Ui32(v151)>>(uint(int32(6))%32))&int32(2))))
											*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)) = uint8(v178)
											F_XLogBeginInsert(m)
											mBase = m.M
											v181 = m.ExcPending
											if v181 != 0 {
												return
											} else {
												F_XLogRegisterData(m, v12+int32(4), int32(8))
												mBase = m.M
												v186 = m.ExcPending
												if v186 != 0 {
													return
												} else {
													F_XLogRegisterBuffer(m, int32(0), v21, int32(8))
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
															*(*int64)(unsafe.Add(mBase, uint32(v40))) = base.I64_rotr(v193, int64(32))
															v200 = int32(4556740)
															v202 = *(*int32)(unsafe.Add(mBase, _consts[14]))
															*(*int32)(unsafe.Add(mBase, _consts[14])) = v202 - int32(1)
															F_LockBuffer(m, v21, int32(0))
															mBase = m.M
															v208 = m.ExcPending
															if v208 != 0 {
																return
															} else {
																v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+20)))
																if v209&int32(4) != 0 {
																	F_heap_toast_delete(m, l0, v12+int32(12), int32(1))
																	mBase = m.M
																	v216 = m.ExcPending
																	if v216 != 0 {
																		return
																	} else {
																		F_ReleaseBuffer(m, v21)
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
																				m.G0 = v12 + int32(32)
																				return
																			}
																		}
																	}
																} else {
																	F_ReleaseBuffer(m, v21)
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
																			m.G0 = v12 + int32(32)
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
									*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)) = uint8(v148)
									v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+18)))
									v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+20)))
									v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)))
									*(*uint16)(unsafe.Add(mBase, uint32(v12)+8)) = uint16(v152)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v14
									v159 = int32(1)
									v163 = int32(4)
									v178 = int32(base.Ui32(v150)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v151)>>(uint(v159)%32))&v148 | (int32(base.Ui32(v151)>>(uint(v163)%32))&v163 | (int32(base.Ui32(v151)>>(uint(int32(12))%32))&v159 | int32(base.Ui32(v151)>>(uint(int32(6))%32))&int32(2))))
									*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)) = uint8(v178)
									F_XLogBeginInsert(m)
									mBase = m.M
									v181 = m.ExcPending
									if v181 != 0 {
										return
									} else {
										F_XLogRegisterData(m, v12+int32(4), int32(8))
										mBase = m.M
										v186 = m.ExcPending
										if v186 != 0 {
											return
										} else {
											F_XLogRegisterBuffer(m, int32(0), v21, int32(8))
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
													*(*int64)(unsafe.Add(mBase, uint32(v40))) = base.I64_rotr(v193, int64(32))
													v200 = int32(4556740)
													v202 = *(*int32)(unsafe.Add(mBase, _consts[14]))
													*(*int32)(unsafe.Add(mBase, _consts[14])) = v202 - int32(1)
													F_LockBuffer(m, v21, int32(0))
													mBase = m.M
													v208 = m.ExcPending
													if v208 != 0 {
														return
													} else {
														v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+20)))
														if v209&int32(4) != 0 {
															F_heap_toast_delete(m, l0, v12+int32(12), int32(1))
															mBase = m.M
															v216 = m.ExcPending
															if v216 != 0 {
																return
															} else {
																F_ReleaseBuffer(m, v21)
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
																		m.G0 = v12 + int32(32)
																		return
																	}
																}
															}
														} else {
															F_ReleaseBuffer(m, v21)
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
																	m.G0 = v12 + int32(32)
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
						F_errmsg_internal(m, int32(269742), int32(0))
						mBase = m.M
						v231 = m.ExcPending
						if v231 != 0 {
							return
						} else {
							F_errfinish(m, int32(522199), int32(6222), int32(360639))
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
					v35 = int32(0)
					if l4 != 0 {
						v38 = F_palloc(m, int32(16))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = v38
							*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v40
							if int32(0) < l2 {
								v46 = F_palloc(m, l2*int32(48))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v48 = v46
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
								v48 = v35
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
						v40 = v35
						*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v40
						if int32(0) < l2 {
							v46 = F_palloc(m, l2*int32(48))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v48 = v46
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
							v48 = v35
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
				v35 = int32(0)
				if l4 != 0 {
					v38 = F_palloc(m, int32(16))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = v38
						*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v40
						if int32(0) < l2 {
							v46 = F_palloc(m, l2*int32(48))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v48 = v46
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
							v48 = v35
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
					v40 = v35
					*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v40
					if int32(0) < l2 {
						v46 = F_palloc(m, l2*int32(48))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v48 = v46
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
						v48 = v35
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
	var v31 int32
	_ = v31
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
	v31 = v4
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
	v38 = v35 + v31*int32(36)
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
	v93 = v31 + int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+112))
	if v93 < v94 {
		v31 = v93
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
	var v22 int32
	_ = v22
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
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
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
	v22 = v4
	goto L4
L4:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+l2))))
	if v28 != 0 {
		v126 = v22
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return v126
L6:
	;
	v130 = v18 + int32(1)
	if v130 != v10 {
		v18 = v130
		v22 = v126
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
	v126 = v121 + v122
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
		v121 = int32(6)
		v122 = v22
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
	v84 = (v22 + v78 - int32(1)) & (v82 - v78)
	if v82 < v36 {
		v121 = v36
		v122 = v84
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
	v126 = v49 + v22
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
	v126 = (v22+v65-int32(1))&(int32(0)-v65) + v73
	goto L6
L23:
	;
	v87 = F_strlen(m, v32)
	mBase = m.M
	v121 = v87 + int32(1)
	v122 = v84
	goto L8
L24:
	;
	v98 = int32(18)
	if v91&int32(255) == v98 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v104 = v98
	goto L27
L26:
	;
	v104 = int32(2)
	goto L27
L27:
	;
	v121 = v104
	v122 = v22
	goto L8
L28:
	;
	v121 = int32(base.Ui32(v55&int32(254)) >> (uint(int32(1)) % 32))
	v122 = v22
	goto L8
L29:
	;
	goto L30
L30:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+12)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v121 = int32(base.Ui32(v118) >> (uint(int32(2)) % 32))
	v122 = (v22 + v111 - int32(1)) & (int32(0) - v111)
	goto L8
L31:
	;
	goto L5
}
func F_heap_copy_minimal_tuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = F_palloc(m, v3+l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v11 = F__emscripten_memset_bulkmem(m, v5, base.I32_extend8_s(int32(0)), l1)
		mBase = m.M
		v12 = v11 + l1
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v13 != 0 {
			v14 = F__emscripten_memcpy_bulkmem(m, v12, l0, v13)
			mBase = m.M
			v15 = v14
		} else {
			v15 = v12
		}
		return v15
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = F_table_slot_create(m, v6, int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[242]))
		v15 = int32(*(*uint8)(unsafe.Add(mBase, _consts[243])))
		if v15&int32(1) != 0 {
			v18 = int32(0)
		} else {
			v18 = v14
		}
		if v18 == int32(0) {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+188))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
			v25 = m.T0[v24].(func(*base.Module, int32, int32, int32, int32) int32)(m, v21, l1, v22, v8)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v8 != 0 {
					F_ExecDropSingleTupleTableSlot(m, v8)
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
				F_errmsg_internal(m, int32(353288), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(343283), int32(1264), int32(284132))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
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
				v25 = *(*int32)(unsafe.Add(mBase, _consts[12]))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v25+(v15^int32(-1))<<(uint(int32(2))%32))))
				v39 = v31
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, _consts[13]))
				v39 = v33 + v15<<(uint(int32(13))%32) + int32(-8192)
			}
			v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
			if v40 == int32(0) {
				v113 = int32(0)
				F_LockBuffer(m, v15, v113)
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return int32(0)
				} else {
					v117 = v113
					F_ReleaseBuffer(m, v15)
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v117
						*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v117
						return v117
					}
				}
			} else {
				v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+12)))
				if base.Ui32(int32(25)) <= base.Ui32(v43) {
					v51 = int32(base.Ui32(v43+int32(262120)) >> (uint(int32(2)) % 32))
				} else {
					v51 = int32(0)
				}
				if base.Ui32(v51&int32(65535)) < base.Ui32(v40) {
					v113 = int32(0)
					F_LockBuffer(m, v15, v113)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						v117 = v113
						F_ReleaseBuffer(m, v15)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v117
							*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v117
							return v117
						}
					}
				} else {
					v59 = v40<<(uint(int32(2))%32) + v39 + int32(20)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
					if v60&int32(98304) != int32(32768) {
						v113 = int32(0)
						F_LockBuffer(m, v15, v113)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							v117 = v113
							F_ReleaseBuffer(m, v15)
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v117
								*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v117
								return v117
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v39 + v60&int32(32767)
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
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
								v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+20)))
								v82 = int32(768)
								if v81&v82 != v82 {
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
									v87 = v86
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
								v99 = int32(0)
								F_HeapCheckForSerializableConflictOut(m, v99, l0, l2, v15, l1)
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int32(0)
								} else {
									F_LockBuffer(m, v15, int32(0))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int32(0)
									} else {
										if l4 == int32(0) {
											v117 = v99
											F_ReleaseBuffer(m, v15)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = v117
												*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v117
												return v117
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
	var v30 int32
	_ = v30
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
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v19 = v16 + int32(12)
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = int32(0)
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
	v226 = m.ExcPending
	if v226 != 0 {
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
	v49 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v41^int32(-1))<<(uint(int32(2))%32))))
	v63 = v55
	goto L6
L8:
	;
	goto L9
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[13]))
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
	v74 = int32(base.Ui32(v66+int32(262120)) >> (uint(int32(2)) % 32))
	goto L14
L13:
	;
	v74 = int32(0)
	goto L14
L14:
	;
	if base.Ui32(v74&int32(65535)) < base.Ui32(v31) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v82 = v31<<(uint(int32(2))%32) + v63 + int32(20)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v83&int32(98304) != int32(32768) {
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
	v94 = v63 + v91&int32(32767)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(base.Ui32(v96) >> (uint(int32(17)) % 32))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v100
	if v30 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+20)))
	v104 = int32(768)
	if v103&v104 != v104 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v114 = F_HeapTupleSatisfiesVisibility(m, v16+int32(8), v23, v41)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L3
	} else {
		goto L24
	}
L20:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v109 = v108
	goto L22
L21:
	;
	v109 = int32(2)
	goto L22
L22:
	;
	if v109 != v30 {
		goto L10
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	F_HeapCheckForSerializableConflictOut(m, v114, v24, v16+int32(8), v41, v23)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	if v114 != 0 {
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
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+21)))
	if v124&int32(8) != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	v127 = F_HeapTupleHeaderIsOnlyLocked(m, v123)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	if v127 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129)+16)))
	if v130 == int32(65533) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129)+12)))
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129)+14)))
	if v133&v134 == int32(65535) {
		goto L10
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v139 = v129 + int32(12)
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+2)))
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19))))
	v142 = int32(16)
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+2)))
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139))))
	if v140|v141<<(uint(v142)%32) == v145|v146<<(uint(v142)%32) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	goto L34
L36:
	;
	if v156 != 0 {
		goto L10
	} else {
		goto L42
	}
L37:
	;
	goto L36
L38:
	;
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+4)))
	if v152 == v153 {
		v156 = int32(1)
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v156 = int32(0)
	goto L37
L41:
	;
	goto L40
L42:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+16)))
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+14)))
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+12)))
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+20)))
	if v162&int32(6272) != int32(4096) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_UnlockReleaseBuffer(m, v41)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L3
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v169 = int32(0)
	v173 = F_GetMultiXactIdMembers(m, v158, v16+int32(28), v169)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L3
	} else {
		goto L47
	}
L46:
	;
	v30 = v158
	v31 = v159
	v32 = v160
	v33 = v161
	goto L1
L47:
	;
	if int32(0) < v173 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v179 = int32(0)
	goto L53
L49:
	;
	v212 = v169
	goto L50
L50:
	;
	F_UnlockReleaseBuffer(m, v41)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L3
	} else {
		goto L58
	}
L51:
	;
	F_pfree(m, v178)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L3
	} else {
		goto L57
	}
L52:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v204 = v202
	goto L51
L53:
	;
	v194 = v178 + v179<<(uint(int32(3))%32)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v195) {
		goto L52
	} else {
		goto L55
	}
L54:
	;
	v204 = int32(0)
	goto L51
L55:
	;
	v199 = v179 + int32(1)
	if v199 != v173 {
		v179 = v199
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v212 = v204
	goto L50
L58:
	;
	v30 = v212
	v31 = v159
	v32 = v160
	v33 = v161
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
			F_errmsg_internal(m, int32(506250), v7)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(523921), int32(761), int32(217039))
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
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
	var v184 int32
	_ = v184
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v300 int32
	_ = v300
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v481 int32
	_ = v481
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v644 int32
	_ = v644
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v685 int64
	_ = v685
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v724 int32
	_ = v724
	var v736 int32
	_ = v736
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v878 int32
	_ = v878
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v977 int32
	_ = v977
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v991 int32
	_ = v991
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1028 int32
	_ = v1028
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1128 int32
	_ = v1128
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1175 int32
	_ = v1175
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
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
	if v81 != 0 {
		goto L340
	} else {
		goto L341
	}
L2:
	;
	v1233 = v1209
	v1234 = v1210
	v1245 = int32(1)
	goto L1
L3:
	;
	v977 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v977)
	if v962&int32(4) != 0 {
		goto L285
	} else {
		goto L286
	}
L4:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v26)+168))
	F_pfree(m, v930)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L44
	} else {
		goto L272
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L44
	} else {
		goto L268
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L44
	} else {
		goto L264
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L44
	} else {
		goto L259
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L44
	} else {
		goto L255
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L44
	} else {
		goto L251
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L44
	} else {
		goto L247
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L44
	} else {
		goto L243
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
	v81 = v6
	v82 = int32(1)
	goto L14
L14:
	;
	if base.Ui32(v80&int32(65535)) < base.Ui32(int32(16384)) {
		v94 = v80
		v95 = v6
		goto L27
	} else {
		goto L28
	}
L15:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	v80 = v79
	v81 = v78
	v82 = v44
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
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v94&int32(4096) != 0 {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v87) < base.Ui32(int32(3)) {
		v94 = v80
		v95 = v6
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v90 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v90)
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	v94 = v93
	v95 = v90
	goto L27
L30:
	;
	if v96 == int32(0) {
		goto L236
	} else {
		goto L237
	}
L31:
	;
	v749 = int32(1)
	v750 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	if v750&int32(128) != 0 {
		v1209 = v6
		v1210 = v749
		goto L2
	} else {
		goto L234
	}
L32:
	;
	v1233 = v736
	v1234 = v6
	v1245 = int32(0)
	goto L1
L33:
	;
	v99 = int32(2)
	if v96 == int32(0) {
		v962 = v99
		v964 = v6
		goto L3
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if base.Ui32(v96) < base.Ui32(int32(3)) {
		goto L30
	} else {
		goto L223
	}
L36:
	;
	if v94&int32(4304) == int32(4224) {
		v962 = v99
		v964 = v6
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	goto L38
L38:
	;
	if int32(base.Ui32(v96-v106)>>(uint(int32(31))%32)) != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	v118 = int32(base.Ui32(v94&int32(128))>>(uint(int32(7))%32)) | base.B2i32(v94&int32(4176) == int32(64))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	goto L40
L40:
	;
	if int32(base.Ui32(v96-v119)>>(uint(int32(31))%32)) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v123 = F_MultiXactIdIsRunning(m, v96, v118)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v254 = F_GetMultiXactIdMembers(m, v96, v26+int32(168), v118)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L44
	} else {
		goto L79
	}
L44:
	;
	return int32(0)
L45:
	;
	if v123 != 0 {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	if v118 != 0 {
		v962 = v99
		v964 = v6
		goto L3
	} else {
		goto L47
	}
L47:
	;
	v130 = F_GetMultiXactIdMembers(m, v96, v26+int32(172), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L44
	} else {
		goto L48
	}
L48:
	;
	if int32(0) < v130 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v26)+172))
	v141 = int32(0)
	goto L54
L50:
	;
	v184 = v6
	goto L51
L51:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v197))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v184)) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L52:
	;
	F_pfree(m, v135)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L44
	} else {
		goto L58
	}
L53:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v171 = v169
	goto L52
L54:
	;
	v161 = v135 + v141<<(uint(int32(3))%32)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v162) {
		goto L53
	} else {
		goto L56
	}
L55:
	;
	v171 = int32(0)
	goto L52
L56:
	;
	v166 = v141 + int32(1)
	if v166 != v130 {
		v141 = v166
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v184 = v171
	goto L51
L59:
	;
	if v209 != 0 {
		goto L8
	} else {
		goto L63
	}
L60:
	;
	v209 = base.B2i32(base.Ui32(v184) < base.Ui32(v197))
	goto L59
L61:
	;
	goto L62
L62:
	;
	v209 = int32(base.Ui32(v184-v197) >> (uint(int32(31)) % 32))
	goto L59
L63:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v210))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v184)) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v222 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v222 = base.B2i32(base.Ui32(v184) < base.Ui32(v210))
	goto L64
L66:
	;
	goto L67
L67:
	;
	v222 = int32(base.Ui32(v184-v210) >> (uint(int32(31)) % 32))
	goto L64
L68:
	;
	v962 = int32(4)
	v964 = v184
	goto L3
L69:
	;
	goto L70
L70:
	;
	v226 = F_TransactionIdDidCommit(m, v184)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L44
	} else {
		goto L71
	}
L71:
	;
	if v226 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v962 = v99
	v964 = int32(0)
	goto L3
L73:
	;
	goto L74
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L44
	} else {
		goto L75
	}
L75:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L44
	} else {
		goto L76
	}
L76:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+88)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v26)+84)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = v96
	F_errmsg_internal(m, int32(53494), v26+int32(80))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L44
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(522199), int32(6789), int32(487224))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L44
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	if v254 <= int32(0) {
		v962 = v99
		v964 = v6
		goto L3
	} else {
		goto L80
	}
L80:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v265 = int32(0)
	v268 = v258
	goto L82
L81:
	;
	v336 = F_palloc(m, v254<<(uint(int32(3))%32))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L44
	} else {
		goto L104
	}
L82:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v26)+168))
	v284 = int32(3)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v283+v265<<(uint(v284)%32))))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v288))&base.B2i32(base.Ui32(v284) <= base.Ui32(v287)) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	goto L97
L84:
	;
	if v300 != 0 {
		goto L81
	} else {
		goto L88
	}
L85:
	;
	v300 = base.B2i32(base.Ui32(v287) < base.Ui32(v288))
	goto L84
L86:
	;
	goto L87
L87:
	;
	v300 = int32(base.Ui32(v287-v288) >> (uint(int32(31)) % 32))
	goto L84
L88:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v268))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v287)) == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	if v312 != 0 {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v312 = base.B2i32(base.Ui32(v287) < base.Ui32(v268))
	goto L89
L91:
	;
	goto L92
L92:
	;
	v312 = int32(base.Ui32(v287-v268) >> (uint(int32(31)) % 32))
	goto L89
L93:
	;
	v313 = v287
	goto L95
L94:
	;
	v313 = v268
	goto L95
L95:
	;
	v315 = v265 + int32(1)
	if v315 != v254 {
		v265 = v315
		v268 = v313
		goto L82
	} else {
		goto L96
	}
L96:
	;
	goto L83
L97:
	;
	if int32(base.Ui32(v96-v317)>>(uint(int32(31))%32)) != 0 {
		goto L81
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v313
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L99
L99:
	;
	if int32(base.Ui32(v96-v322)>>(uint(int32(31))%32)) != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v96
	goto L102
L101:
	;
	goto L102
L102:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v26)+168))
	F_pfree(m, v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L44
	} else {
		goto L103
	}
L103:
	;
	v736 = int32(0)
	goto L32
L104:
	;
	v347 = int32(0)
	v349 = v6
	v350 = int32(0)
	v351 = v6
	v352 = v6
	goto L105
L105:
	;
	v362 = int32(3)
	v363 = v347 << (uint(v362) % 32)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v26)+168))
	v365 = v363 + v364
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	if base.Ui32(v367) <= base.Ui32(v362) {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	goto L4
L107:
	;
	v695 = v347 + int32(1)
	if v254 != v695 {
		v347 = v695
		v349 = v690
		v350 = v691
		v351 = v692
		v352 = v693
		goto L105
	} else {
		goto L222
	}
L108:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v26)+168))
	v685 = *(*int64)(unsafe.Add(mBase, uint32(v683+v363)))
	*(*int64)(unsafe.Add(mBase, uint32(v336+v350<<(uint(int32(3))%32)))) = v685
	v690 = v677
	v691 = v350 + int32(1)
	v692 = v678
	v693 = v679
	goto L107
L109:
	;
	if base.Ui32(v366) < base.Ui32(int32(3)) {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	goto L111
L111:
	;
	if v349 != 0 {
		goto L7
	} else {
		goto L166
	}
L112:
	;
	if v489 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L113:
	;
	v489 = int32(0)
	goto L112
L114:
	;
	goto L115
L115:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v380 == v366 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v489 = int32(1)
	goto L112
L117:
	;
	goto L118
L118:
	;
	v384 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	if v384 <= int32(0) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v489 = v481
	goto L112
L120:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v388 == int32(0) {
		v481 = int32(0)
		goto L119
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v450 = *(*int32)(unsafe.Add(mBase, _consts[41]))
	v452 = int32(0)
	v454 = v384 - int32(1)
	goto L142
L123:
	;
	v393 = v388
	goto L124
L124:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v393)+20))
	if v398 == int32(4) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v481 = int32(0)
	goto L119
L126:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v393)+80))
	if v445 != 0 {
		v393 = v445
		goto L124
	} else {
		goto L141
	}
L127:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	if v401 == int32(0) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v404 = int32(1)
	if v366 == v401 {
		v481 = v404
		goto L119
	} else {
		goto L129
	}
L129:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v393)+52))
	v408 = v406 - int32(1)
	if v408 < int32(0) {
		goto L126
	} else {
		goto L130
	}
L130:
	;
	v413 = int32(0)
	v415 = v408
	goto L131
L131:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v393)+48))
	v421 = int32(2)
	v422 = base.I32_div_s(v415-v413, v421)
	v423 = v422 + v413
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v419+v423<<(uint(v421)%32))))
	if v427 == v366 {
		v481 = v404
		goto L119
	} else {
		goto L133
	}
L132:
	;
	goto L126
L133:
	;
	v431 = F_TransactionIdPrecedes(m, v427, v366)
	mBase = m.M
	if v431 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v432 = v423 + int32(1)
	goto L136
L135:
	;
	v432 = v413
	goto L136
L136:
	;
	if v431 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v435 = v415
	goto L139
L138:
	;
	v435 = v423 - int32(1)
	goto L139
L139:
	;
	if v432 <= v435 {
		v413 = v432
		v415 = v435
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
	v459 = int32(2)
	v460 = base.I32_div_s(v454-v452, v459)
	v461 = v460 + v452
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v450+v461<<(uint(v459)%32))))
	v466 = base.B2i32(v465 == v366)
	if v465 == v366 {
		v481 = v466
		goto L119
	} else {
		goto L144
	}
L143:
	;
	v481 = v466
	goto L119
L144:
	;
	v469 = base.B2i32(base.Ui32(v465) < base.Ui32(v366))
	if base.Ui32(v465) < base.Ui32(v366) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v470 = v461 + int32(1)
	goto L147
L146:
	;
	v470 = v452
	goto L147
L147:
	;
	if base.Ui32(v465) < base.Ui32(v366) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v473 = v454
	goto L150
L149:
	;
	v473 = v461 - int32(1)
	goto L150
L150:
	;
	if v470 <= v473 {
		v452 = v470
		v454 = v473
		goto L142
	} else {
		goto L151
	}
L151:
	;
	goto L143
L152:
	;
	v492 = F_TransactionIdIsInProgress(m, v366)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L44
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v497))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v366)) == int32(0) {
		goto L158
	} else {
		goto L159
	}
L155:
	;
	if v492 == int32(0) {
		v690 = v349
		v691 = v350
		v692 = v351
		v693 = v352
		goto L107
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	if v509 == int32(0) {
		v677 = v349
		v678 = v351
		v679 = int32(1)
		goto L108
	} else {
		goto L161
	}
L158:
	;
	v509 = base.B2i32(base.Ui32(v366) < base.Ui32(v497))
	goto L157
L159:
	;
	goto L160
L160:
	;
	v509 = int32(base.Ui32(v366-v497) >> (uint(int32(31)) % 32))
	goto L157
L161:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L44
	} else {
		goto L162
	}
L162:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L44
	} else {
		goto L163
	}
L163:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+152)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v26)+148)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v26)+144)) = v96
	F_errmsg_internal(m, int32(53418), v26+int32(144))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L44
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(522199), int32(6910), int32(487224))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L44
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	if base.Ui32(v366) < base.Ui32(int32(3)) {
		goto L169
	} else {
		goto L170
	}
L167:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v663))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v366)) == int32(0) {
		goto L218
	} else {
		goto L219
	}
L168:
	;
	if v652 != 0 {
		goto L208
	} else {
		goto L209
	}
L169:
	;
	v652 = int32(0)
	goto L168
L170:
	;
	goto L171
L171:
	;
	v543 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v543 == v366 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v652 = int32(1)
	goto L168
L173:
	;
	goto L174
L174:
	;
	v547 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	if v547 <= int32(0) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v652 = v644
	goto L168
L176:
	;
	v551 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v551 == int32(0) {
		v644 = int32(0)
		goto L175
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v613 = *(*int32)(unsafe.Add(mBase, _consts[41]))
	v615 = int32(0)
	v617 = v547 - int32(1)
	goto L198
L179:
	;
	v556 = v551
	goto L180
L180:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v556)+20))
	if v561 == int32(4) {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v644 = int32(0)
	goto L175
L182:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v556)+80))
	if v608 != 0 {
		v556 = v608
		goto L180
	} else {
		goto L197
	}
L183:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	if v564 == int32(0) {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v567 = int32(1)
	if v366 == v564 {
		v644 = v567
		goto L175
	} else {
		goto L185
	}
L185:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v556)+52))
	v571 = v569 - int32(1)
	if v571 < int32(0) {
		goto L182
	} else {
		goto L186
	}
L186:
	;
	v576 = int32(0)
	v578 = v571
	goto L187
L187:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v556)+48))
	v584 = int32(2)
	v585 = base.I32_div_s(v578-v576, v584)
	v586 = v585 + v576
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v582+v586<<(uint(v584)%32))))
	if v590 == v366 {
		v644 = v567
		goto L175
	} else {
		goto L189
	}
L188:
	;
	goto L182
L189:
	;
	v594 = F_TransactionIdPrecedes(m, v590, v366)
	mBase = m.M
	if v594 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v595 = v586 + int32(1)
	goto L192
L191:
	;
	v595 = v576
	goto L192
L192:
	;
	if v594 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v598 = v578
	goto L195
L194:
	;
	v598 = v586 - int32(1)
	goto L195
L195:
	;
	if v595 <= v598 {
		v576 = v595
		v578 = v598
		goto L187
	} else {
		goto L196
	}
L196:
	;
	goto L188
L197:
	;
	goto L181
L198:
	;
	v622 = int32(2)
	v623 = base.I32_div_s(v617-v615, v622)
	v624 = v623 + v615
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v613+v624<<(uint(v622)%32))))
	v629 = base.B2i32(v628 == v366)
	if v628 == v366 {
		v644 = v629
		goto L175
	} else {
		goto L200
	}
L199:
	;
	v644 = v629
	goto L175
L200:
	;
	v632 = base.B2i32(base.Ui32(v628) < base.Ui32(v366))
	if base.Ui32(v628) < base.Ui32(v366) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v633 = v624 + int32(1)
	goto L203
L202:
	;
	v633 = v615
	goto L203
L203:
	;
	if base.Ui32(v628) < base.Ui32(v366) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v636 = v617
	goto L206
L205:
	;
	v636 = v624 - int32(1)
	goto L206
L206:
	;
	if v633 <= v636 {
		v615 = v633
		v617 = v636
		goto L198
	} else {
		goto L207
	}
L207:
	;
	goto L199
L208:
	;
	v661 = v351
	goto L167
L209:
	;
	goto L210
L210:
	;
	v653 = F_TransactionIdIsInProgress(m, v366)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L44
	} else {
		goto L211
	}
L211:
	;
	if v653 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v661 = v351
	goto L167
L213:
	;
	goto L214
L214:
	;
	v657 = F_TransactionIdDidCommit(m, v366)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L44
	} else {
		goto L215
	}
L215:
	;
	if v657 == int32(0) {
		v690 = int32(0)
		v691 = v350
		v692 = v351
		v693 = v352
		goto L107
	} else {
		goto L216
	}
L216:
	;
	v661 = int32(1)
	goto L167
L217:
	;
	if v675 != 0 {
		goto L6
	} else {
		goto L221
	}
L218:
	;
	v675 = base.B2i32(base.Ui32(v366) < base.Ui32(v663))
	goto L217
L219:
	;
	goto L220
L220:
	;
	v675 = int32(base.Ui32(v366-v663) >> (uint(int32(31)) % 32))
	goto L217
L221:
	;
	v677 = v366
	v678 = v661
	v679 = v352
	goto L108
L222:
	;
	goto L106
L223:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v699))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v96)) == int32(0) {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	if v711 != 0 {
		goto L5
	} else {
		goto L228
	}
L225:
	;
	v711 = base.B2i32(base.Ui32(v96) < base.Ui32(v699))
	goto L224
L226:
	;
	goto L227
L227:
	;
	v711 = int32(base.Ui32(v96-v699) >> (uint(int32(31)) % 32))
	goto L224
L228:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v712))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v96)) == int32(0) {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	if v724 != 0 {
		goto L31
	} else {
		goto L233
	}
L230:
	;
	v724 = base.B2i32(base.Ui32(v96) < base.Ui32(v712))
	goto L229
L231:
	;
	goto L232
L232:
	;
	v724 = int32(base.Ui32(v96-v712) >> (uint(int32(31)) % 32))
	goto L229
L233:
	;
	v736 = v6
	goto L32
L234:
	;
	if v750&int32(4176) == int32(64) {
		v1209 = v6
		v1210 = v749
		goto L2
	} else {
		goto L235
	}
L235:
	;
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)))
	v759 = v757 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v759)
	v1209 = v6
	v1210 = v749
	goto L2
L236:
	;
	v1233 = int32(1)
	v1234 = v6
	v1245 = int32(0)
	goto L1
L237:
	;
	goto L238
L238:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L44
	} else {
		goto L239
	}
L239:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L44
	} else {
		goto L240
	}
L240:
	;
	v772 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v96
	F_errmsg_internal(m, int32(335602), v26+int32(16))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L44
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(522199), int32(7258), int32(402345))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L44
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L44
	} else {
		goto L244
	}
L244:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+164)) = v792
	*(*int32)(unsafe.Add(mBase, uint32(v26)+160)) = v42
	F_errmsg_internal(m, int32(58629), v26+int32(160))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L44
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(522199), int32(7096), int32(402345))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L44
	} else {
		goto L246
	}
L246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L247:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L44
	} else {
		goto L248
	}
L248:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v812
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v96
	F_errmsg_internal(m, int32(58735), v26+int32(32))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L44
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(522199), int32(6744), int32(487224))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L44
	} else {
		goto L250
	}
L250:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L251:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L44
	} else {
		goto L252
	}
L252:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v832
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v96
	F_errmsg_internal(m, int32(352048), v26+int32(48))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L44
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(522199), int32(6760), int32(487224))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L44
	} else {
		goto L254
	}
L254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L255:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L44
	} else {
		goto L256
	}
L256:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = v852
	*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v96
	F_errmsg_internal(m, int32(58671), v26-int32(-64))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L44
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(522199), int32(6776), int32(487224))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L44
	} else {
		goto L258
	}
L258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L259:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L44
	} else {
		goto L260
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+128)) = v96
	F_errmsg_internal(m, int32(144570), v26+int32(128))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L44
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+116)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v26)+112)) = v349
	F_errdetail_internal(m, int32(605448), v26+int32(112))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L44
	} else {
		goto L262
	}
L262:
	;
	F_errfinish(m, int32(522199), int32(6935), int32(487224))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L44
	} else {
		goto L263
	}
L263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L264:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L44
	} else {
		goto L265
	}
L265:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+104)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v26)+100)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v96
	F_errmsg_internal(m, int32(53494), v26+int32(96))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L44
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(522199), int32(6972), int32(487224))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L44
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L44
	} else {
		goto L269
	}
L269:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v96
	F_errmsg_internal(m, int32(58587), v26)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L44
	} else {
		goto L270
	}
L270:
	;
	F_errfinish(m, int32(522199), int32(7235), int32(402345))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L44
	} else {
		goto L271
	}
L271:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L272:
	;
	if v691 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	F_pfree(m, v336)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L44
	} else {
		goto L284
	}
L274:
	;
	v950 = int32(2)
	v951 = int32(0)
	goto L273
L275:
	;
	goto L276
L276:
	;
	v937 = int32(0)
	if v693|base.B2i32(v690 == v937) == v937 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	if v692&int32(1) != 0 {
		goto L280
	} else {
		goto L281
	}
L278:
	;
	goto L279
L279:
	;
	v948 = F_MultiXactIdCreateFromMembers(m, v691, v336)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L44
	} else {
		goto L283
	}
L280:
	;
	v946 = int32(20)
	goto L282
L281:
	;
	v946 = int32(4)
	goto L282
L282:
	;
	v950 = v946
	v951 = v690
	goto L273
L283:
	;
	v950 = int32(8)
	v951 = v948
	goto L273
L284:
	;
	v962 = v950
	v964 = v951
	goto L3
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v964
	v983 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
	v985 = v983 & int32(-7377)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v985)
	if v962&int32(16) != 0 {
		goto L288
	} else {
		goto L289
	}
L286:
	;
	goto L287
L287:
	;
	if v962&int32(8) == int32(0) {
		goto L291
	} else {
		goto L292
	}
L288:
	;
	v991 = v985 | int32(1024)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v991)
	goto L290
L289:
	;
	goto L290
L290:
	;
	v1233 = int32(0)
	v1234 = v977
	v1245 = int32(0)
	goto L1
L291:
	;
	v1209 = int32(0)
	v1210 = v977
	goto L2
L292:
	;
	goto L293
L293:
	;
	v999 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
	v1001 = v999 & int32(58159)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v1001)
	v1003 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	v1005 = v1003 & int32(57343)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)) = uint16(v1005)
	v1007 = int32(0)
	v1011 = F_GetMultiXactIdMembers(m, v964, v26+int32(172), v1007)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L44
	} else {
		goto L295
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v964
	v1191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
	v1192 = v1191 | v1189
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v1192)
	v1194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	v1195 = v1194 | v1175
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)) = uint16(v1195)
	v1233 = v1007
	v1234 = v977
	v1245 = int32(0)
	goto L1
L295:
	;
	if v1011 <= int32(0) {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	v1175 = int32(0)
	v1189 = int32(4240)
	goto L294
L297:
	;
	goto L298
L298:
	;
	v1017 = int32(1)
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v26)+172))
	if v1011 == v1017 {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	if v1011&v1017 == int32(0) {
		v1144 = v1105
		v1147 = v1109
		v1148 = v1113
		goto L320
	} else {
		goto L321
	}
L300:
	;
	v1022 = int32(0)
	v1105 = v1022
	v1108 = v1022
	v1109 = v1022
	v1113 = v1022
	goto L299
L301:
	;
	goto L302
L302:
	;
	v1028 = int32(0)
	v1038 = v1028
	v1040 = v1028
	v1041 = v1028
	v1042 = v1028
	v1046 = v1028
	goto L303
L303:
	;
	v1058 = v1019 + v1041<<(uint(int32(3))%32)
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+4))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1059<<(uint(int32(2))%32))+uint32(_consts[37])))
	if base.Ui32(v1038) < base.Ui32(v1064) {
		goto L305
	} else {
		goto L306
	}
L304:
	;
	v1105 = v1094
	v1108 = v1096
	v1109 = v1092
	v1113 = v1093
	goto L299
L305:
	;
	v1066 = v1064
	goto L307
L306:
	;
	v1066 = v1038
	goto L307
L307:
	;
	switch v1059 - int32(3) {
	case 0:
		goto L311
	case 1:
		v1073 = v1042
		goto L309
	case 2:
		goto L310
	default:
		v1075 = v1042
		v1076 = v1046
		goto L308
	}
L308:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+12))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1077<<(uint(int32(2))%32))+uint32(_consts[37])))
	switch v1077 - int32(3) {
	case 0:
		goto L315
	case 1:
		v1090 = v1075
		goto L313
	case 2:
		goto L314
	default:
		v1092 = v1075
		v1093 = v1076
		goto L312
	}
L309:
	;
	v1075 = v1073
	v1076 = int32(1)
	goto L308
L310:
	;
	v1073 = v1042 | int32(8192)
	goto L309
L311:
	;
	v1075 = v1042 | int32(8192)
	v1076 = v1046
	goto L308
L312:
	;
	if base.Ui32(v1066) < base.Ui32(v1082) {
		goto L316
	} else {
		goto L317
	}
L313:
	;
	v1092 = v1090
	v1093 = int32(1)
	goto L312
L314:
	;
	v1090 = v1075 | int32(8192)
	goto L313
L315:
	;
	v1092 = v1075 | int32(8192)
	v1093 = v1076
	goto L312
L316:
	;
	v1094 = v1082
	goto L318
L317:
	;
	v1094 = v1066
	goto L318
L318:
	;
	v1095 = int32(2)
	v1096 = v1041 + v1095
	v1098 = v1040 + v1095
	if v1098 != v1011&int32(2147483646) {
		v1038 = v1094
		v1040 = v1098
		v1041 = v1096
		v1042 = v1092
		v1046 = v1093
		goto L303
	} else {
		goto L319
	}
L319:
	;
	goto L304
L320:
	;
	F_pfree(m, v1019)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L44
	} else {
		goto L328
	}
L321:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1019+v1108<<(uint(int32(3))%32))+4))
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1128<<(uint(int32(2))%32))+uint32(_consts[37])))
	if base.Ui32(v1105) < base.Ui32(v1133) {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	v1135 = v1133
	goto L324
L323:
	;
	v1135 = v1105
	goto L324
L324:
	;
	switch v1128 - int32(3) {
	case 0:
		goto L327
	case 1:
		v1142 = v1109
		goto L325
	case 2:
		goto L326
	default:
		v1144 = v1135
		v1147 = v1109
		v1148 = v1113
		goto L320
	}
L325:
	;
	v1144 = v1135
	v1147 = v1142
	v1148 = int32(1)
	goto L320
L326:
	;
	v1142 = v1109 | int32(8192)
	goto L325
L327:
	;
	v1144 = v1135
	v1147 = v1109 | int32(8192)
	v1148 = v1113
	goto L320
L328:
	;
	if v1144&int32(-2) == int32(2) {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	if v1148 != 0 {
		v1175 = v1147
		v1189 = int32(4160)
		goto L294
	} else {
		goto L332
	}
L330:
	;
	goto L331
L331:
	;
	if v1144 != 0 {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	v1175 = v1147
	v1189 = int32(4288)
	goto L294
L333:
	;
	v1160 = int32(4096)
	goto L335
L334:
	;
	v1160 = int32(4112)
	goto L335
L335:
	;
	if v1144 == int32(1) {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v1163 = int32(4176)
	goto L338
L337:
	;
	v1163 = v1160
	goto L338
L338:
	;
	if v1148 != 0 {
		v1175 = v1147
		v1189 = v1163
		goto L294
	} else {
		goto L339
	}
L339:
	;
	v1175 = v1147
	v1189 = v1163 | int32(128)
	goto L294
L340:
	;
	v1246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
	v1248 = v1246 | int32(768)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v1248)
	goto L342
L341:
	;
	goto L342
L342:
	;
	if v95 != 0 {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v1250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)))
	v1253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	if v1253&int32(16384) != 0 {
		goto L346
	} else {
		goto L347
	}
L344:
	;
	goto L345
L345:
	;
	if v1245 != 0 {
		goto L349
	} else {
		goto L350
	}
L346:
	;
	v1256 = int32(4)
	goto L348
L347:
	;
	v1256 = int32(2)
	goto L348
L348:
	;
	v1257 = v1250 | v1256
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v1257)
	goto L345
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	v1261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	v1263 = v1261 & int32(40959)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)) = uint16(v1263)
	v1265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
	v1269 = v1265&int32(58159) | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)) = uint16(v1269)
	goto L351
L350:
	;
	goto L351
L351:
	;
	v1273 = (v81 | v82) & (v1245 | v1233)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v1273)
	if v1233&v82 != 0 {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	m.G0 = v26 + int32(176)
	return v1234 | v95 | v81
L353:
	;
	v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v1276&int32(1) != 0 {
		goto L352
	} else {
		goto L354
	}
L354:
	;
	v1283 = F_heap_tuple_should_freeze(m, l0, l1, l2+int32(12), l2+int32(16))
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L44
	} else {
		goto L355
	}
L355:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v1283)
	goto L352
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
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v227 int32
	_ = v227
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v297 int32
	_ = v297
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v367 int32
	_ = v367
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v426 int32
	_ = v426
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
	v31 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31+(v22^int32(-1))<<(uint(int32(2))%32))))
	v45 = v37
	goto L4
L6:
	;
	goto L7
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v45 = v39 + v22<<(uint(int32(13))%32) + int32(-8192)
	goto L4
L8:
	;
	v56 = int32(base.Ui32(v46+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v410
	F_LockBuffer(m, v22, int32(0))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
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
	v202 = int32(base.Ui32(v20) >> (uint(int32(16)) % 32))
	v206 = int32(1)
	v209 = int32(0)
	v210 = v206
	v212 = v206
	goto L48
L24:
	;
	if v88&v87 == int32(0) {
		v410 = v169
		goto L11
	} else {
		goto L46
	}
L25:
	;
	v102 = int32(0)
	v106 = v102
	v109 = v102
	v110 = int32(1)
	v111 = v78
	goto L37
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
		goto L36
	}
L29:
	;
	v410 = int32(0)
	goto L11
L30:
	;
	goto L31
L31:
	;
	v78 = int32(1)
	v79 = int32(2)
	v83 = (v56 + v78) & int32(65535)
	if base.Ui32(v83) <= base.Ui32(v79) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v86 = v79
	goto L34
L33:
	;
	v86 = v83
	goto L34
L34:
	;
	v87 = int32(1)
	v88 = v86 - v87
	v92 = l0 + int32(108)
	v94 = v45 + int32(24)
	if base.Ui32(int32(3)) <= base.Ui32(v83) {
		goto L25
	} else {
		goto L35
	}
L35:
	;
	v97 = int32(0)
	v169 = v97
	v174 = v78
	v182 = v97
	goto L24
L36:
	;
	v410 = int32(0)
	goto L11
L37:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v110<<(uint(int32(2))%32)+v94-int32(4))))
	if v124&int32(98304) == int32(32768) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v169 = v156
	v174 = v158
	v182 = v160 - int32(1)
	goto L24
L39:
	;
	v129 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v92+v106<<(uint(v129)%32)))) = uint16(v111)
	v135 = v106 + v129
	goto L41
L40:
	;
	v135 = v106
	goto L41
L41:
	;
	v137 = v111 + int32(1)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v137&int32(65535)<<(uint(int32(2))%32)+v94-int32(4))))
	if v145&int32(98304) == int32(32768) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v150 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v92+v135<<(uint(v150)%32)))) = uint16(v137)
	v156 = v135 + v150
	goto L44
L43:
	;
	v156 = v135
	goto L44
L44:
	;
	v157 = int32(2)
	v158 = v111 + v157
	v159 = int32(65535)
	v160 = v158 & v159
	v162 = v109 + v157
	if v162&v159 != v88&int32(65534) {
		v106 = v156
		v109 = v162
		v110 = v160
		v111 = v158
		goto L37
	} else {
		goto L45
	}
L45:
	;
	goto L38
L46:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v94+v182<<(uint(int32(2))%32))))
	if v188&int32(98304) != int32(32768) {
		v410 = v169
		goto L11
	} else {
		goto L47
	}
L47:
	;
	v193 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v92+v169<<(uint(v193)%32)))) = uint16(v174)
	v410 = v169 + v193
	goto L11
L48:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v210<<(uint(int32(2))%32)+(v45+int32(24))-int32(4))))
	if v227&int32(98304) == int32(32768) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v410 = v256
	goto L11
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(base.Ui32(v227) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v45 + v227&int32(32767)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+20)) = uint16(v212)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+18)) = uint16(v20)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)) = uint16(v202)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v240
	F_HeapCheckForSerializableConflictOut(m, int32(1), v239, v17+int32(12), v22, v19)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	v256 = v209
	goto L52
L52:
	;
	v260 = v212 + int32(1)
	v262 = v260 & int32(65535)
	if base.Ui32(v262) <= base.Ui32(v56) {
		v209 = v256
		v210 = v262
		v212 = v260
		goto L48
	} else {
		goto L54
	}
L53:
	;
	v250 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(108)+v209<<(uint(v250)%32)))) = uint16(v212)
	v256 = v209 + v250
	goto L52
L54:
	;
	goto L49
L55:
	;
	goto L13
L56:
	;
	v410 = int32(0)
	goto L11
L57:
	;
	goto L58
L58:
	;
	v272 = int32(base.Ui32(v20) >> (uint(int32(16)) % 32))
	v276 = int32(1)
	v279 = int32(0)
	v280 = v276
	v282 = v276
	goto L59
L59:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v280<<(uint(int32(2))%32)+(v45+int32(24))-int32(4))))
	if v297&int32(98304) != int32(32768) {
		v327 = v279
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v410 = v327
	goto L11
L61:
	;
	v330 = v282 + int32(1)
	v332 = v330 & int32(65535)
	if base.Ui32(v332) <= base.Ui32(v56) {
		v279 = v327
		v280 = v332
		v282 = v330
		goto L59
	} else {
		goto L65
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(base.Ui32(v297) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v45 + v297&int32(32767)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+20)) = uint16(v282)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+18)) = uint16(v20)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)) = uint16(v272)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v310
	v317 = F_HeapTupleSatisfiesVisibility(m, v17+int32(12), v19, v22)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v317 == int32(0) {
		v327 = v279
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v321 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(108)+v279<<(uint(v321)%32)))) = uint16(v282)
	v327 = v279 + v321
	goto L61
L65:
	;
	goto L60
L66:
	;
	v410 = int32(0)
	goto L11
L67:
	;
	goto L68
L68:
	;
	v342 = int32(base.Ui32(v20) >> (uint(int32(16)) % 32))
	v346 = int32(1)
	v349 = int32(0)
	v350 = v346
	v352 = v346
	goto L69
L69:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v350<<(uint(int32(2))%32)+(v45+int32(24))-int32(4))))
	if v367&int32(98304) != int32(32768) {
		v402 = v349
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v410 = v402
	goto L11
L71:
	;
	v405 = v352 + int32(1)
	v407 = v405 & int32(65535)
	if base.Ui32(v407) <= base.Ui32(v56) {
		v349 = v402
		v350 = v407
		v352 = v405
		goto L69
	} else {
		goto L76
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(base.Ui32(v367) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v45 + v367&int32(32767)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+20)) = uint16(v352)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+18)) = uint16(v20)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)) = uint16(v342)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v380
	v387 = F_HeapTupleSatisfiesVisibility(m, v17+int32(12), v19, v22)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_HeapCheckForSerializableConflictOut(m, v387, v389, v17+int32(12), v22, v19)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if v387 == int32(0) {
		v402 = v349
		goto L71
	} else {
		goto L75
	}
L75:
	;
	v396 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(108)+v349<<(uint(v396)%32)))) = uint16(v352)
	v402 = v349 + v396
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
