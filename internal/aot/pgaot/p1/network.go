package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_network_callback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int64
	_ = v133
	var v135 int32
	_ = v135
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	v6 = m.G0
	v8 = v6 - int32(128)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v10 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v13 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v8 + int32(128)
	return
L4:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v299)
	goto L3
L5:
	;
	v16 = int32(0)
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v23 = m.G0
	v25 = v23 - int32(32)
	m.G0 = v25
	goto L11
L6:
	;
	goto L7
L7:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229))))
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v230 != v231 {
		goto L47
	} else {
		goto L48
	}
L8:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160))))
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v161 != v162 {
		v299 = v16
		goto L4
	} else {
		goto L36
	}
L9:
	;
	m.G0 = v25 + int32(32)
	goto L8
L10:
	;
	switch v18 - int32(2) {
	case 0:
		goto L21
	default:
		goto L9
	case 8:
		goto L20
	}
L11:
	;
	if v18 == int32(2) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v33 = int32(32)
	goto L16
L15:
	;
	v33 = int32(128)
	goto L16
L16:
	;
	goto L10
L19:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v8))) = uint16(v18)
	goto L9
L20:
	;
	if base.Ui32(int32(128)) < base.Ui32(v33) {
		goto L9
	} else {
		goto L26
	}
L21:
	;
	if base.Ui32(int32(32)) < base.Ui32(v33) {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v53
	v58 = int32(-1) << (uint(int32(32)-v33) % 32)
	v59 = int32(24)
	v61 = int32(65280)
	v63 = int32(8)
	if v33 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v75 = v58<<(uint(v59)%32) | v58&v61<<(uint(v63)%32) | (int32(base.Ui32(v58)>>(uint(v63)%32))&v61 | int32(base.Ui32(v58)>>(uint(v59)%32)))
	goto L25
L24:
	;
	v75 = v53
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v75
	goto L19
L26:
	;
	v79 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v79
	v82 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v82
	v85 = v25 + int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v85))) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v82
	v91 = v79
	v93 = v33
	goto L27
L27:
	;
	v98 = int32(0)
	if v93 <= v98 {
		v108 = v98
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v135
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v137
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v139
	goto L19
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v91+v85))) = uint8(v108)
	v113 = int32(0)
	v115 = v93 - int32(8)
	if v115 <= v113 {
		v125 = v113
		goto L32
	} else {
		goto L33
	}
L30:
	;
	if base.Ui32(int32(7)) < base.Ui32(v93) {
		v108 = int32(255)
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v108 = int32(255) << (uint(int32(8)-v93) % 32)
	goto L29
L32:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v85+(v91|int32(1))))) = uint8(v125)
	v127 = int32(16)
	v130 = v91 + int32(2)
	if v130 != v127 {
		v91 = v130
		v93 = v93 - v127
		goto L27
	} else {
		goto L35
	}
L33:
	;
	if base.Ui32(int32(7)) < base.Ui32(v115) {
		v125 = int32(255)
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v125 = int32(255) << (uint(int32(16)-v93) % 32)
	goto L32
L35:
	;
	goto L28
L36:
	;
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160))))
	switch v167 - int32(2) {
	case 0:
		goto L40
	default:
		goto L38
	case 8:
		goto L39
	}
L37:
	;
	if v225 == int32(0) {
		v299 = v16
		goto L4
	} else {
		goto L46
	}
L38:
	;
	v225 = int32(0)
	goto L37
L39:
	;
	v177 = int32(8)
	v178 = v8 + v177
	v180 = l0 + v177
	v182 = v160 + v177
	v184 = int32(0)
	goto L41
L40:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	v225 = base.B2i32(v170&(v171^v172) == int32(0))
	goto L37
L41:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184+v178))))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184+v180))))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184+v182))))
	if v191&(v193^v195) != 0 {
		goto L38
	} else {
		goto L43
	}
L42:
	;
	v225 = int32(1)
	goto L37
L43:
	;
	v199 = v184 | int32(1)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180+v199))))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v199))))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199+v178))))
	if (v201^v203)&v206 != 0 {
		goto L38
	} else {
		goto L44
	}
L44:
	;
	v209 = v184 + int32(2)
	if v209 != int32(16) {
		v184 = v209
		goto L41
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	v299 = int32(1)
	goto L4
L47:
	;
	v299 = int32(0)
	goto L4
L48:
	;
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229))))
	switch v236 - int32(2) {
	case 0:
		goto L52
	default:
		goto L50
	case 8:
		goto L51
	}
L49:
	;
	if v294 == int32(0) {
		goto L47
	} else {
		goto L58
	}
L50:
	;
	v294 = int32(0)
	goto L49
L51:
	;
	v246 = int32(8)
	v247 = l1 + v246
	v249 = l0 + v246
	v251 = v229 + v246
	v253 = int32(0)
	goto L53
L52:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	v294 = base.B2i32(v239&(v240^v241) == int32(0))
	goto L49
L53:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253+v247))))
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253+v249))))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253+v251))))
	if v260&(v262^v264) != 0 {
		goto L50
	} else {
		goto L55
	}
L54:
	;
	v294 = int32(1)
	goto L49
L55:
	;
	v268 = v253 | int32(1)
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249+v268))))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251+v268))))
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268+v247))))
	if (v270^v272)&v275 != 0 {
		goto L50
	} else {
		goto L56
	}
L56:
	;
	v278 = v253 + int32(2)
	if v278 != int32(16) {
		v253 = v278
		goto L53
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	v299 = int32(1)
	goto L4
}
func F_network_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v17 = int32(1)
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
			if v19&v17 != 0 {
				v22 = v17
			} else {
				v22 = int32(4)
			}
			v23 = v3 + v22
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
			v25 = int32(1)
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v27&v25 != 0 {
				v30 = v25
			} else {
				v30 = int32(4)
			}
			v31 = v8 + v30
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
			if v24 == v32 {
				v34 = int32(2)
				v35 = v23 + v34
				v37 = v31 + v34
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
				if base.Ui32(v38) < base.Ui32(v39) {
					v41 = v38
				} else {
					v41 = v39
				}
				v43 = int32(base.Ui32(v41) >> (uint(int32(3)) % 32))
				v44 = F_memcmp(m, v35, v37, v43)
				mBase = m.M
				if v44 != 0 {
					v136 = v44
					v155 = v136
				} else {
					v46 = v41 & int32(7)
					if v46 == int32(0) {
						v127 = v38 - v39
						if v127 != 0 {
							v136 = v127
							v155 = v136
						} else {
							if v24 == int32(2) {
								v132 = int32(4)
							} else {
								v132 = int32(16)
							}
							v133 = F_memcmp(m, v35, v37, v132)
							mBase = m.M
							v155 = v133
						}
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v35))))
						v51 = int32(128)
						v52 = v50 & v51
						v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v37))))
						if v52 != v54&v51 {
							v143 = v52
							if v143 != 0 {
								v146 = int32(1)
							} else {
								v146 = int32(-1)
							}
							v155 = v146
						} else {
							if v46 == int32(1) {
								v127 = v38 - v39
								if v127 != 0 {
									v136 = v127
									v155 = v136
								} else {
									if v24 == int32(2) {
										v132 = int32(4)
									} else {
										v132 = int32(16)
									}
									v133 = F_memcmp(m, v35, v37, v132)
									mBase = m.M
									v155 = v133
								}
							} else {
								v60 = int32(1)
								v62 = int32(128)
								v63 = v50 << (uint(v60) % 32) & v62
								if v63 != v54<<(uint(v60)%32)&v62 {
									v143 = v63
									if v143 != 0 {
										v146 = int32(1)
									} else {
										v146 = int32(-1)
									}
									v155 = v146
								} else {
									if base.Ui32(v46) < base.Ui32(int32(3)) {
										v127 = v38 - v39
										if v127 != 0 {
											v136 = v127
											v155 = v136
										} else {
											if v24 == int32(2) {
												v132 = int32(4)
											} else {
												v132 = int32(16)
											}
											v133 = F_memcmp(m, v35, v37, v132)
											mBase = m.M
											v155 = v133
										}
									} else {
										v71 = int32(2)
										v73 = int32(128)
										v74 = v50 << (uint(v71) % 32) & v73
										if v74 != v54<<(uint(v71)%32)&v73 {
											v143 = v74
											if v143 != 0 {
												v146 = int32(1)
											} else {
												v146 = int32(-1)
											}
											v155 = v146
										} else {
											if v46 == int32(3) {
												v127 = v38 - v39
												if v127 != 0 {
													v136 = v127
													v155 = v136
												} else {
													if v24 == int32(2) {
														v132 = int32(4)
													} else {
														v132 = int32(16)
													}
													v133 = F_memcmp(m, v35, v37, v132)
													mBase = m.M
													v155 = v133
												}
											} else {
												v82 = int32(3)
												v84 = int32(128)
												v85 = v50 << (uint(v82) % 32) & v84
												if v85 != v54<<(uint(v82)%32)&v84 {
													v143 = v85
													if v143 != 0 {
														v146 = int32(1)
													} else {
														v146 = int32(-1)
													}
													v155 = v146
												} else {
													if base.Ui32(v46) < base.Ui32(int32(5)) {
														v127 = v38 - v39
														if v127 != 0 {
															v136 = v127
															v155 = v136
														} else {
															if v24 == int32(2) {
																v132 = int32(4)
															} else {
																v132 = int32(16)
															}
															v133 = F_memcmp(m, v35, v37, v132)
															mBase = m.M
															v155 = v133
														}
													} else {
														v93 = int32(4)
														v95 = int32(128)
														v96 = v50 << (uint(v93) % 32) & v95
														if v96 != v54<<(uint(v93)%32)&v95 {
															v143 = v96
															if v143 != 0 {
																v146 = int32(1)
															} else {
																v146 = int32(-1)
															}
															v155 = v146
														} else {
															if v46 == int32(5) {
																v127 = v38 - v39
																if v127 != 0 {
																	v136 = v127
																	v155 = v136
																} else {
																	if v24 == int32(2) {
																		v132 = int32(4)
																	} else {
																		v132 = int32(16)
																	}
																	v133 = F_memcmp(m, v35, v37, v132)
																	mBase = m.M
																	v155 = v133
																}
															} else {
																v104 = int32(5)
																v106 = int32(128)
																v107 = v50 << (uint(v104) % 32) & v106
																if v107 != v54<<(uint(v104)%32)&v106 {
																	v143 = v107
																	if v143 != 0 {
																		v146 = int32(1)
																	} else {
																		v146 = int32(-1)
																	}
																	v155 = v146
																} else {
																	if v46 != int32(7) {
																		v127 = v38 - v39
																		if v127 != 0 {
																			v136 = v127
																			v155 = v136
																		} else {
																			if v24 == int32(2) {
																				v132 = int32(4)
																			} else {
																				v132 = int32(16)
																			}
																			v133 = F_memcmp(m, v35, v37, v132)
																			mBase = m.M
																			v155 = v133
																		}
																	} else {
																		v115 = int32(6)
																		v117 = int32(128)
																		v118 = v50 << (uint(v115) % 32) & v117
																		if v118 != v54<<(uint(v115)%32)&v117 {
																			v143 = v118
																			if v143 != 0 {
																				v146 = int32(1)
																			} else {
																				v146 = int32(-1)
																			}
																			v155 = v146
																		} else {
																			v127 = v38 - v39
																			if v127 != 0 {
																				v136 = v127
																				v155 = v136
																			} else {
																				if v24 == int32(2) {
																					v132 = int32(4)
																				} else {
																					v132 = int32(16)
																				}
																				v133 = F_memcmp(m, v35, v37, v132)
																				mBase = m.M
																				v155 = v133
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
							}
						}
					}
				}
			} else {
				v136 = v24 - v32
				v155 = v136
			}
			return v155
		}
	}
}
func F_network_cmp_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	v10 = int32(1)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v12&v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if v260 != 0 {
		goto L71
	} else {
		goto L72
	}
L2:
	;
	return v252
L3:
	;
	v15 = v10
	goto L5
L4:
	;
	v15 = int32(4)
	goto L5
L5:
	;
	v16 = l0 + v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v18 = int32(1)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v20&v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v23 = v18
	goto L8
L7:
	;
	v23 = int32(4)
	goto L8
L8:
	;
	v24 = l1 + v23
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v17 == v25 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v27 = int32(2)
	v28 = v16 + v27
	v30 = v24 + v27
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if base.Ui32(v31) < base.Ui32(v32) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v252 = v17 - v25
	goto L2
L12:
	;
	v34 = v31
	goto L14
L13:
	;
	v34 = v32
	goto L14
L14:
	;
	v36 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(4)) <= base.Ui32(v36) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	if v98 != 0 {
		v252 = v98
		goto L2
	} else {
		goto L33
	}
L16:
	;
	v98 = int32(0)
	goto L15
L17:
	;
	v72 = v67
	v73 = v68
	v74 = v69
	goto L27
L18:
	;
	if (v28|v30)&int32(3) != 0 {
		v67 = v28
		v68 = v30
		v69 = v36
		goto L17
	} else {
		goto L21
	}
L19:
	;
	v60 = v28
	v61 = v30
	v62 = v36
	goto L20
L20:
	;
	if v62 == int32(0) {
		goto L16
	} else {
		goto L26
	}
L21:
	;
	v44 = v28
	v45 = v30
	v46 = v36
	goto L22
L22:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v49 != v50 {
		v67 = v44
		v68 = v45
		v69 = v46
		goto L17
	} else {
		goto L24
	}
L23:
	;
	v60 = v55
	v61 = v53
	v62 = v57
	goto L20
L24:
	;
	v52 = int32(4)
	v53 = v45 + v52
	v55 = v44 + v52
	v57 = v46 - v52
	if base.Ui32(int32(3)) < base.Ui32(v57) {
		v44 = v55
		v45 = v53
		v46 = v57
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v67 = v60
	v68 = v61
	v69 = v62
	goto L17
L27:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v77 == v78 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v98 = v77 - v78
	goto L15
L29:
	;
	v80 = int32(1)
	v85 = v74 - v80
	if v85 != 0 {
		v72 = v72 + v80
		v73 = v73 + v80
		v74 = v85
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	goto L16
L33:
	;
	v100 = v34 & int32(7)
	if v100 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v181 = v31 - v32
	if v181 != 0 {
		v252 = v181
		goto L2
	} else {
		goto L49
	}
L35:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v28))))
	v105 = int32(128)
	v106 = v104 & v105
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v30))))
	if v106 != v108&v105 {
		v260 = v106
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v100 == int32(1) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v114 = int32(1)
	v116 = int32(128)
	v117 = v104 << (uint(v114) % 32) & v116
	if v117 != v108<<(uint(v114)%32)&v116 {
		v260 = v117
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if base.Ui32(v100) < base.Ui32(int32(3)) {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v125 = int32(2)
	v127 = int32(128)
	v128 = v104 << (uint(v125) % 32) & v127
	if v128 != v108<<(uint(v125)%32)&v127 {
		v260 = v128
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v100 == int32(3) {
		goto L34
	} else {
		goto L41
	}
L41:
	;
	v136 = int32(3)
	v138 = int32(128)
	v139 = v104 << (uint(v136) % 32) & v138
	if v139 != v108<<(uint(v136)%32)&v138 {
		v260 = v139
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if base.Ui32(v100) < base.Ui32(int32(5)) {
		goto L34
	} else {
		goto L43
	}
L43:
	;
	v147 = int32(4)
	v149 = int32(128)
	v150 = v104 << (uint(v147) % 32) & v149
	if v150 != v108<<(uint(v147)%32)&v149 {
		v260 = v150
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v100 == int32(5) {
		goto L34
	} else {
		goto L45
	}
L45:
	;
	v158 = int32(5)
	v160 = int32(128)
	v161 = v104 << (uint(v158) % 32) & v160
	if v161 != v108<<(uint(v158)%32)&v160 {
		v260 = v161
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v100 != int32(7) {
		goto L34
	} else {
		goto L47
	}
L47:
	;
	v169 = int32(6)
	v171 = int32(128)
	v172 = v104 << (uint(v169) % 32) & v171
	if v172 != v108<<(uint(v169)%32)&v171 {
		v260 = v172
		goto L1
	} else {
		goto L48
	}
L48:
	;
	goto L34
L49:
	;
	if v17 == int32(2) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v186 = int32(4)
	goto L52
L51:
	;
	v186 = int32(16)
	goto L52
L52:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v186) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	return v248
L54:
	;
	v248 = int32(0)
	goto L53
L55:
	;
	v222 = v217
	v223 = v218
	v224 = v219
	goto L65
L56:
	;
	if (v28|v30)&int32(3) != 0 {
		v217 = v28
		v218 = v30
		v219 = v186
		goto L55
	} else {
		goto L59
	}
L57:
	;
	v210 = v28
	v211 = v30
	v212 = v186
	goto L58
L58:
	;
	if v212 == int32(0) {
		goto L54
	} else {
		goto L64
	}
L59:
	;
	v194 = v28
	v195 = v30
	v196 = v186
	goto L60
L60:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	if v199 != v200 {
		v217 = v194
		v218 = v195
		v219 = v196
		goto L55
	} else {
		goto L62
	}
L61:
	;
	v210 = v205
	v211 = v203
	v212 = v207
	goto L58
L62:
	;
	v202 = int32(4)
	v203 = v195 + v202
	v205 = v194 + v202
	v207 = v196 - v202
	if base.Ui32(int32(3)) < base.Ui32(v207) {
		v194 = v205
		v195 = v203
		v196 = v207
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v217 = v210
	v218 = v211
	v219 = v212
	goto L55
L65:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	if v227 == v228 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v248 = v227 - v228
	goto L53
L67:
	;
	v230 = int32(1)
	v235 = v224 - v230
	if v235 != 0 {
		v222 = v222 + v230
		v223 = v223 + v230
		v224 = v235
		goto L65
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	goto L66
L70:
	;
	goto L54
L71:
	;
	v263 = int32(1)
	goto L73
L72:
	;
	v263 = int32(-1)
	goto L73
L73:
	;
	return v263
}
func F_network_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v17 = int32(1)
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
			if v19&v17 != 0 {
				v22 = v17
			} else {
				v22 = int32(4)
			}
			v23 = v3 + v22
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
			v25 = int32(1)
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v27&v25 != 0 {
				v30 = v25
			} else {
				v30 = int32(4)
			}
			v31 = v8 + v30
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
			if v24 == v32 {
				v34 = int32(2)
				v35 = v23 + v34
				v37 = v31 + v34
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
				if base.Ui32(v38) < base.Ui32(v39) {
					v41 = v38
				} else {
					v41 = v39
				}
				v43 = int32(base.Ui32(v41) >> (uint(int32(3)) % 32))
				v44 = F_memcmp(m, v35, v37, v43)
				mBase = m.M
				if v44 != 0 {
					v136 = v44
					v155 = v136
				} else {
					v46 = v41 & int32(7)
					if v46 == int32(0) {
						v127 = v38 - v39
						if v127 != 0 {
							v136 = v127
							v155 = v136
						} else {
							if v24 == int32(2) {
								v132 = int32(4)
							} else {
								v132 = int32(16)
							}
							v133 = F_memcmp(m, v35, v37, v132)
							mBase = m.M
							v155 = v133
						}
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v35))))
						v51 = int32(128)
						v52 = v50 & v51
						v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v37))))
						if v52 != v54&v51 {
							v143 = v52
							if v143 != 0 {
								v146 = int32(1)
							} else {
								v146 = int32(-1)
							}
							v155 = v146
						} else {
							if v46 == int32(1) {
								v127 = v38 - v39
								if v127 != 0 {
									v136 = v127
									v155 = v136
								} else {
									if v24 == int32(2) {
										v132 = int32(4)
									} else {
										v132 = int32(16)
									}
									v133 = F_memcmp(m, v35, v37, v132)
									mBase = m.M
									v155 = v133
								}
							} else {
								v60 = int32(1)
								v62 = int32(128)
								v63 = v50 << (uint(v60) % 32) & v62
								if v63 != v54<<(uint(v60)%32)&v62 {
									v143 = v63
									if v143 != 0 {
										v146 = int32(1)
									} else {
										v146 = int32(-1)
									}
									v155 = v146
								} else {
									if base.Ui32(v46) < base.Ui32(int32(3)) {
										v127 = v38 - v39
										if v127 != 0 {
											v136 = v127
											v155 = v136
										} else {
											if v24 == int32(2) {
												v132 = int32(4)
											} else {
												v132 = int32(16)
											}
											v133 = F_memcmp(m, v35, v37, v132)
											mBase = m.M
											v155 = v133
										}
									} else {
										v71 = int32(2)
										v73 = int32(128)
										v74 = v50 << (uint(v71) % 32) & v73
										if v74 != v54<<(uint(v71)%32)&v73 {
											v143 = v74
											if v143 != 0 {
												v146 = int32(1)
											} else {
												v146 = int32(-1)
											}
											v155 = v146
										} else {
											if v46 == int32(3) {
												v127 = v38 - v39
												if v127 != 0 {
													v136 = v127
													v155 = v136
												} else {
													if v24 == int32(2) {
														v132 = int32(4)
													} else {
														v132 = int32(16)
													}
													v133 = F_memcmp(m, v35, v37, v132)
													mBase = m.M
													v155 = v133
												}
											} else {
												v82 = int32(3)
												v84 = int32(128)
												v85 = v50 << (uint(v82) % 32) & v84
												if v85 != v54<<(uint(v82)%32)&v84 {
													v143 = v85
													if v143 != 0 {
														v146 = int32(1)
													} else {
														v146 = int32(-1)
													}
													v155 = v146
												} else {
													if base.Ui32(v46) < base.Ui32(int32(5)) {
														v127 = v38 - v39
														if v127 != 0 {
															v136 = v127
															v155 = v136
														} else {
															if v24 == int32(2) {
																v132 = int32(4)
															} else {
																v132 = int32(16)
															}
															v133 = F_memcmp(m, v35, v37, v132)
															mBase = m.M
															v155 = v133
														}
													} else {
														v93 = int32(4)
														v95 = int32(128)
														v96 = v50 << (uint(v93) % 32) & v95
														if v96 != v54<<(uint(v93)%32)&v95 {
															v143 = v96
															if v143 != 0 {
																v146 = int32(1)
															} else {
																v146 = int32(-1)
															}
															v155 = v146
														} else {
															if v46 == int32(5) {
																v127 = v38 - v39
																if v127 != 0 {
																	v136 = v127
																	v155 = v136
																} else {
																	if v24 == int32(2) {
																		v132 = int32(4)
																	} else {
																		v132 = int32(16)
																	}
																	v133 = F_memcmp(m, v35, v37, v132)
																	mBase = m.M
																	v155 = v133
																}
															} else {
																v104 = int32(5)
																v106 = int32(128)
																v107 = v50 << (uint(v104) % 32) & v106
																if v107 != v54<<(uint(v104)%32)&v106 {
																	v143 = v107
																	if v143 != 0 {
																		v146 = int32(1)
																	} else {
																		v146 = int32(-1)
																	}
																	v155 = v146
																} else {
																	if v46 != int32(7) {
																		v127 = v38 - v39
																		if v127 != 0 {
																			v136 = v127
																			v155 = v136
																		} else {
																			if v24 == int32(2) {
																				v132 = int32(4)
																			} else {
																				v132 = int32(16)
																			}
																			v133 = F_memcmp(m, v35, v37, v132)
																			mBase = m.M
																			v155 = v133
																		}
																	} else {
																		v115 = int32(6)
																		v117 = int32(128)
																		v118 = v50 << (uint(v115) % 32) & v117
																		if v118 != v54<<(uint(v115)%32)&v117 {
																			v143 = v118
																			if v143 != 0 {
																				v146 = int32(1)
																			} else {
																				v146 = int32(-1)
																			}
																			v155 = v146
																		} else {
																			v127 = v38 - v39
																			if v127 != 0 {
																				v136 = v127
																				v155 = v136
																			} else {
																				if v24 == int32(2) {
																					v132 = int32(4)
																				} else {
																					v132 = int32(16)
																				}
																				v133 = F_memcmp(m, v35, v37, v132)
																				mBase = m.M
																				v155 = v133
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
							}
						}
					}
				}
			} else {
				v136 = v24 - v32
				v155 = v136
			}
			return base.B2i32(v155 == int32(0))
		}
	}
}
func F_network_family(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = int32(1)
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
		if v12&v10 != 0 {
			v15 = v10
		} else {
			v15 = int32(4)
		}
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+v15))))
		if v17 == int32(3) {
			v20 = int32(6)
		} else {
			v20 = int32(0)
		}
		if v17 == int32(2) {
			v23 = int32(4)
		} else {
			v23 = v20
		}
		return v23
	}
}
func F_network_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v17 = int32(1)
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
			if v19&v17 != 0 {
				v22 = v17
			} else {
				v22 = int32(4)
			}
			v23 = v3 + v22
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
			v25 = int32(1)
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v27&v25 != 0 {
				v30 = v25
			} else {
				v30 = int32(4)
			}
			v31 = v8 + v30
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
			if v24 == v32 {
				v34 = int32(2)
				v35 = v23 + v34
				v37 = v31 + v34
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
				if base.Ui32(v38) < base.Ui32(v39) {
					v41 = v38
				} else {
					v41 = v39
				}
				v43 = int32(base.Ui32(v41) >> (uint(int32(3)) % 32))
				v44 = F_memcmp(m, v35, v37, v43)
				mBase = m.M
				if v44 != 0 {
					v136 = v44
					v155 = v136
				} else {
					v46 = v41 & int32(7)
					if v46 == int32(0) {
						v127 = v38 - v39
						if v127 != 0 {
							v136 = v127
							v155 = v136
						} else {
							if v24 == int32(2) {
								v132 = int32(4)
							} else {
								v132 = int32(16)
							}
							v133 = F_memcmp(m, v35, v37, v132)
							mBase = m.M
							v155 = v133
						}
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v35))))
						v51 = int32(128)
						v52 = v50 & v51
						v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v37))))
						if v52 != v54&v51 {
							v143 = v52
							if v143 != 0 {
								v146 = int32(1)
							} else {
								v146 = int32(-1)
							}
							v155 = v146
						} else {
							if v46 == int32(1) {
								v127 = v38 - v39
								if v127 != 0 {
									v136 = v127
									v155 = v136
								} else {
									if v24 == int32(2) {
										v132 = int32(4)
									} else {
										v132 = int32(16)
									}
									v133 = F_memcmp(m, v35, v37, v132)
									mBase = m.M
									v155 = v133
								}
							} else {
								v60 = int32(1)
								v62 = int32(128)
								v63 = v50 << (uint(v60) % 32) & v62
								if v63 != v54<<(uint(v60)%32)&v62 {
									v143 = v63
									if v143 != 0 {
										v146 = int32(1)
									} else {
										v146 = int32(-1)
									}
									v155 = v146
								} else {
									if base.Ui32(v46) < base.Ui32(int32(3)) {
										v127 = v38 - v39
										if v127 != 0 {
											v136 = v127
											v155 = v136
										} else {
											if v24 == int32(2) {
												v132 = int32(4)
											} else {
												v132 = int32(16)
											}
											v133 = F_memcmp(m, v35, v37, v132)
											mBase = m.M
											v155 = v133
										}
									} else {
										v71 = int32(2)
										v73 = int32(128)
										v74 = v50 << (uint(v71) % 32) & v73
										if v74 != v54<<(uint(v71)%32)&v73 {
											v143 = v74
											if v143 != 0 {
												v146 = int32(1)
											} else {
												v146 = int32(-1)
											}
											v155 = v146
										} else {
											if v46 == int32(3) {
												v127 = v38 - v39
												if v127 != 0 {
													v136 = v127
													v155 = v136
												} else {
													if v24 == int32(2) {
														v132 = int32(4)
													} else {
														v132 = int32(16)
													}
													v133 = F_memcmp(m, v35, v37, v132)
													mBase = m.M
													v155 = v133
												}
											} else {
												v82 = int32(3)
												v84 = int32(128)
												v85 = v50 << (uint(v82) % 32) & v84
												if v85 != v54<<(uint(v82)%32)&v84 {
													v143 = v85
													if v143 != 0 {
														v146 = int32(1)
													} else {
														v146 = int32(-1)
													}
													v155 = v146
												} else {
													if base.Ui32(v46) < base.Ui32(int32(5)) {
														v127 = v38 - v39
														if v127 != 0 {
															v136 = v127
															v155 = v136
														} else {
															if v24 == int32(2) {
																v132 = int32(4)
															} else {
																v132 = int32(16)
															}
															v133 = F_memcmp(m, v35, v37, v132)
															mBase = m.M
															v155 = v133
														}
													} else {
														v93 = int32(4)
														v95 = int32(128)
														v96 = v50 << (uint(v93) % 32) & v95
														if v96 != v54<<(uint(v93)%32)&v95 {
															v143 = v96
															if v143 != 0 {
																v146 = int32(1)
															} else {
																v146 = int32(-1)
															}
															v155 = v146
														} else {
															if v46 == int32(5) {
																v127 = v38 - v39
																if v127 != 0 {
																	v136 = v127
																	v155 = v136
																} else {
																	if v24 == int32(2) {
																		v132 = int32(4)
																	} else {
																		v132 = int32(16)
																	}
																	v133 = F_memcmp(m, v35, v37, v132)
																	mBase = m.M
																	v155 = v133
																}
															} else {
																v104 = int32(5)
																v106 = int32(128)
																v107 = v50 << (uint(v104) % 32) & v106
																if v107 != v54<<(uint(v104)%32)&v106 {
																	v143 = v107
																	if v143 != 0 {
																		v146 = int32(1)
																	} else {
																		v146 = int32(-1)
																	}
																	v155 = v146
																} else {
																	if v46 != int32(7) {
																		v127 = v38 - v39
																		if v127 != 0 {
																			v136 = v127
																			v155 = v136
																		} else {
																			if v24 == int32(2) {
																				v132 = int32(4)
																			} else {
																				v132 = int32(16)
																			}
																			v133 = F_memcmp(m, v35, v37, v132)
																			mBase = m.M
																			v155 = v133
																		}
																	} else {
																		v115 = int32(6)
																		v117 = int32(128)
																		v118 = v50 << (uint(v115) % 32) & v117
																		if v118 != v54<<(uint(v115)%32)&v117 {
																			v143 = v118
																			if v143 != 0 {
																				v146 = int32(1)
																			} else {
																				v146 = int32(-1)
																			}
																			v155 = v146
																		} else {
																			v127 = v38 - v39
																			if v127 != 0 {
																				v136 = v127
																				v155 = v136
																			} else {
																				if v24 == int32(2) {
																					v132 = int32(4)
																				} else {
																					v132 = int32(16)
																				}
																				v133 = F_memcmp(m, v35, v37, v132)
																				mBase = m.M
																				v155 = v133
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
							}
						}
					}
				}
			} else {
				v136 = v24 - v32
				v155 = v136
			}
			return int32(base.Ui32(v155^int32(-1)) >> (uint(int32(31)) % 32))
		}
	}
}
func F_network_host(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(1)
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		if v15&v13 != 0 {
			v18 = v13
		} else {
			v18 = int32(4)
		}
		v19 = v9 + v18
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
		v21 = int32(2)
		if v20 == v21 {
			v27 = int32(32)
		} else {
			v27 = int32(128)
		}
		v28 = F_pg_inet_net_ntop(m, v20, v19+v21, v27, v6)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			if v28 != 0 {
				v30 = int32(47)
				v31 = F___strchrnul(m, v6, v30)
				mBase = m.M
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
				if v33 == v30 {
					v37 = v31
				} else {
					v37 = int32(0)
				}
				if v37 != 0 {
					v38 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v38)
				} else {
				}
				v40 = F_cstring_to_text(m, v6)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 - int32(-64)
					return v40
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50462850))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(289854), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(489960), int32(1149), int32(67327))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
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
	}
}
func F_network_network(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
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
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v17 = F_palloc0(m, int32(22))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(4)
			v20 = v17 + v19
			v21 = int32(1)
			v22 = v17 + v21
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			v25 = v23 & v21
			v27 = v12 + v21
			v29 = v12 + v19
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			v32 = v30 & v21
			if v32 != 0 {
				v33 = v27
			} else {
				v33 = v29
			}
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
			if v34 != 0 {
				if v25 != 0 {
					v35 = v22
				} else {
					v35 = v20
				}
				v36 = int32(2)
				v41 = v34
				v42 = int32(0)
				for {
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+(v33+v36)))))
					if base.Ui32(int32(7)) < base.Ui32(v41) {
						v61 = int32(-1)
					} else {
						v61 = int32(255) << (uint(int32(8)-v41) % 32)
					}
					v62 = v53 & v61
					*(*uint8)(unsafe.Add(mBase, uint32(v42+(v35+v36)))) = uint8(v62)
					v66 = int32(8)
					if v41 <= v66 {
						v69 = v66
					} else {
						v69 = v41
					}
					v71 = v69 - int32(8)
					if v71 != 0 {
						v41 = v71
						v42 = v42 + int32(1)
						continue
					} else {
						break
					}
					break
				}
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
				v73 = int32(1)
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				v79 = v75 & v73
				v82 = v72 & v73
			} else {
				v79 = v25
				v82 = v32
			}
			if v79 != 0 {
				v90 = int32(1)
			} else {
				v90 = int32(4)
			}
			if v82 != 0 {
				v94 = int32(1)
			} else {
				v94 = int32(4)
			}
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v94))))
			*(*uint8)(unsafe.Add(mBase, uint32(v17+v90))) = uint8(v96)
			if v79 != 0 {
				v98 = v22
			} else {
				v98 = v20
			}
			v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v99&int32(1) != 0 {
				v102 = v27
			} else {
				v102 = v29
			}
			v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
			*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)) = uint8(v103)
			v107 = int32(1)
			v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			if v109&v107 != 0 {
				v112 = v107
			} else {
				v112 = int32(4)
			}
			v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v112))))
			if v114 == int32(2) {
				v117 = int32(40)
			} else {
				v117 = int32(88)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v117
			return v17
		}
	}
}
func F_network_send(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	v2 = l1
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	F_pq_begintypsend(m, v10)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = int32(1)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v18&v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = v16
	goto L5
L4:
	;
	v21 = int32(4)
	goto L5
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v21))))
	F_enlargeStringInfo(m, v10, int32(1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v27+v28))) = uint8(v23)
	v31 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v27 + v31
	v35 = l0 + v31
	v37 = l0 + int32(4)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v38&v31 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v41 = v35
	goto L9
L8:
	;
	v41 = v37
	goto L9
L9:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	F_enlargeStringInfo(m, v10, int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v46+v47))) = uint8(v42)
	v50 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v46 + v50
	F_enlargeStringInfo(m, v10, v50)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v56+v57))) = uint8(v2)
	v60 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v56 + v60
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v65&v60 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v68 = v60
	goto L14
L13:
	;
	v68 = int32(4)
	goto L14
L14:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v68))))
	F_enlargeStringInfo(m, v10, int32(1))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v70 == int32(2) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v81 = int32(4)
	goto L18
L17:
	;
	v81 = int32(16)
	goto L18
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v74+v75))) = uint8(v81)
	v83 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v74 + v83
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v86&v83 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v89 = v35
	goto L21
L20:
	;
	v89 = v37
	goto L21
L21:
	;
	v93 = int32(0)
	goto L22
L22:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+(v89+int32(2))))))
	F_enlargeStringInfo(m, v10, int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v117 << (uint(int32(2)) % 32)
	goto L26
L24:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v105+v106))) = uint8(v101)
	v109 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v105 + v109
	v113 = v93 + v109
	if v113 != v81 {
		v93 = v113
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	m.G0 = v10 + int32(16)
	return v116
}
func F_network_smaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = F_pg_detoast_datum_packed(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v18 = int32(1)
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
			if v20&v18 != 0 {
				v23 = v18
			} else {
				v23 = int32(4)
			}
			v24 = v4 + v23
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
			v26 = int32(1)
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v28&v26 != 0 {
				v31 = v26
			} else {
				v31 = int32(4)
			}
			v32 = v9 + v31
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
			if v25 == v33 {
				v35 = int32(2)
				v36 = v24 + v35
				v38 = v32 + v35
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
				if base.Ui32(v39) < base.Ui32(v40) {
					v42 = v39
				} else {
					v42 = v40
				}
				v44 = int32(base.Ui32(v42) >> (uint(int32(3)) % 32))
				v45 = F_memcmp(m, v36, v38, v44)
				mBase = m.M
				if v45 != 0 {
					v137 = v45
					v156 = v137
				} else {
					v47 = v42 & int32(7)
					if v47 == int32(0) {
						v128 = v39 - v40
						if v128 != 0 {
							v137 = v128
							v156 = v137
						} else {
							if v25 == int32(2) {
								v133 = int32(4)
							} else {
								v133 = int32(16)
							}
							v134 = F_memcmp(m, v36, v38, v133)
							mBase = m.M
							v156 = v134
						}
					} else {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v36))))
						v52 = int32(128)
						v53 = v51 & v52
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v38))))
						if v53 != v55&v52 {
							v144 = v53
							if v144 != 0 {
								v147 = int32(1)
							} else {
								v147 = int32(-1)
							}
							v156 = v147
						} else {
							if v47 == int32(1) {
								v128 = v39 - v40
								if v128 != 0 {
									v137 = v128
									v156 = v137
								} else {
									if v25 == int32(2) {
										v133 = int32(4)
									} else {
										v133 = int32(16)
									}
									v134 = F_memcmp(m, v36, v38, v133)
									mBase = m.M
									v156 = v134
								}
							} else {
								v61 = int32(1)
								v63 = int32(128)
								v64 = v51 << (uint(v61) % 32) & v63
								if v64 != v55<<(uint(v61)%32)&v63 {
									v144 = v64
									if v144 != 0 {
										v147 = int32(1)
									} else {
										v147 = int32(-1)
									}
									v156 = v147
								} else {
									if base.Ui32(v47) < base.Ui32(int32(3)) {
										v128 = v39 - v40
										if v128 != 0 {
											v137 = v128
											v156 = v137
										} else {
											if v25 == int32(2) {
												v133 = int32(4)
											} else {
												v133 = int32(16)
											}
											v134 = F_memcmp(m, v36, v38, v133)
											mBase = m.M
											v156 = v134
										}
									} else {
										v72 = int32(2)
										v74 = int32(128)
										v75 = v51 << (uint(v72) % 32) & v74
										if v75 != v55<<(uint(v72)%32)&v74 {
											v144 = v75
											if v144 != 0 {
												v147 = int32(1)
											} else {
												v147 = int32(-1)
											}
											v156 = v147
										} else {
											if v47 == int32(3) {
												v128 = v39 - v40
												if v128 != 0 {
													v137 = v128
													v156 = v137
												} else {
													if v25 == int32(2) {
														v133 = int32(4)
													} else {
														v133 = int32(16)
													}
													v134 = F_memcmp(m, v36, v38, v133)
													mBase = m.M
													v156 = v134
												}
											} else {
												v83 = int32(3)
												v85 = int32(128)
												v86 = v51 << (uint(v83) % 32) & v85
												if v86 != v55<<(uint(v83)%32)&v85 {
													v144 = v86
													if v144 != 0 {
														v147 = int32(1)
													} else {
														v147 = int32(-1)
													}
													v156 = v147
												} else {
													if base.Ui32(v47) < base.Ui32(int32(5)) {
														v128 = v39 - v40
														if v128 != 0 {
															v137 = v128
															v156 = v137
														} else {
															if v25 == int32(2) {
																v133 = int32(4)
															} else {
																v133 = int32(16)
															}
															v134 = F_memcmp(m, v36, v38, v133)
															mBase = m.M
															v156 = v134
														}
													} else {
														v94 = int32(4)
														v96 = int32(128)
														v97 = v51 << (uint(v94) % 32) & v96
														if v97 != v55<<(uint(v94)%32)&v96 {
															v144 = v97
															if v144 != 0 {
																v147 = int32(1)
															} else {
																v147 = int32(-1)
															}
															v156 = v147
														} else {
															if v47 == int32(5) {
																v128 = v39 - v40
																if v128 != 0 {
																	v137 = v128
																	v156 = v137
																} else {
																	if v25 == int32(2) {
																		v133 = int32(4)
																	} else {
																		v133 = int32(16)
																	}
																	v134 = F_memcmp(m, v36, v38, v133)
																	mBase = m.M
																	v156 = v134
																}
															} else {
																v105 = int32(5)
																v107 = int32(128)
																v108 = v51 << (uint(v105) % 32) & v107
																if v108 != v55<<(uint(v105)%32)&v107 {
																	v144 = v108
																	if v144 != 0 {
																		v147 = int32(1)
																	} else {
																		v147 = int32(-1)
																	}
																	v156 = v147
																} else {
																	if v47 != int32(7) {
																		v128 = v39 - v40
																		if v128 != 0 {
																			v137 = v128
																			v156 = v137
																		} else {
																			if v25 == int32(2) {
																				v133 = int32(4)
																			} else {
																				v133 = int32(16)
																			}
																			v134 = F_memcmp(m, v36, v38, v133)
																			mBase = m.M
																			v156 = v134
																		}
																	} else {
																		v116 = int32(6)
																		v118 = int32(128)
																		v119 = v51 << (uint(v116) % 32) & v118
																		if v119 != v55<<(uint(v116)%32)&v118 {
																			v144 = v119
																			if v144 != 0 {
																				v147 = int32(1)
																			} else {
																				v147 = int32(-1)
																			}
																			v156 = v147
																		} else {
																			v128 = v39 - v40
																			if v128 != 0 {
																				v137 = v128
																				v156 = v137
																			} else {
																				if v25 == int32(2) {
																					v133 = int32(4)
																				} else {
																					v133 = int32(16)
																				}
																				v134 = F_memcmp(m, v36, v38, v133)
																				mBase = m.M
																				v156 = v134
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
							}
						}
					}
				}
			} else {
				v137 = v25 - v33
				v156 = v137
			}
			if v156 < int32(0) {
				v159 = v4
			} else {
				v159 = v9
			}
			return v159
		}
	}
}
func F_network_sub(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v169 int32
	_ = v169
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = int32(1)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v15&v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = v13
	goto L6
L5:
	;
	v18 = int32(4)
	goto L6
L6:
	;
	v19 = v6 + v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v21 = int32(1)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v23&v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = v21
	goto L9
L8:
	;
	v26 = int32(4)
	goto L9
L9:
	;
	v27 = v11 + v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v20 != v28 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L12
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if base.Ui32(v33) <= base.Ui32(v32) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	goto L15
L15:
	;
	v37 = int32(2)
	v38 = v19 + v37
	v40 = v27 + v37
	v42 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(4)) <= base.Ui32(v42) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if v104 != 0 {
		goto L34
	} else {
		goto L35
	}
L17:
	;
	v104 = int32(0)
	goto L16
L18:
	;
	v78 = v73
	v79 = v74
	v80 = v75
	goto L28
L19:
	;
	if (v38|v40)&int32(3) != 0 {
		v73 = v38
		v74 = v40
		v75 = v42
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v66 = v38
	v67 = v40
	v68 = v42
	goto L21
L21:
	;
	if v68 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L22:
	;
	v50 = v38
	v51 = v40
	v52 = v42
	goto L23
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v55 != v56 {
		v73 = v50
		v74 = v51
		v75 = v52
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v66 = v61
	v67 = v59
	v68 = v63
	goto L21
L25:
	;
	v58 = int32(4)
	v59 = v51 + v58
	v61 = v50 + v58
	v63 = v52 - v58
	if base.Ui32(int32(3)) < base.Ui32(v63) {
		v50 = v61
		v51 = v59
		v52 = v63
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v73 = v66
	v74 = v67
	v75 = v68
	goto L18
L28:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v83 == v84 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v104 = v83 - v84
	goto L16
L30:
	;
	v86 = int32(1)
	v91 = v80 - v86
	if v91 != 0 {
		v78 = v78 + v86
		v79 = v79 + v86
		v80 = v91
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L17
L34:
	;
	return int32(0)
L35:
	;
	goto L36
L36:
	;
	v108 = v32 & int32(7)
	if v108 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	return int32(1)
L38:
	;
	goto L39
L39:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v38))))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v40))))
	v117 = v114 ^ v116
	if base.Ui32(int32(127)) < base.Ui32(v117) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	return int32(0)
L41:
	;
	goto L42
L42:
	;
	v122 = int32(1)
	if v108 == v122 {
		v169 = v122
		goto L43
	} else {
		goto L44
	}
L43:
	;
	return v169
L44:
	;
	if v117<<(uint(int32(1))%32)&int32(128) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	return int32(0)
L46:
	;
	goto L47
L47:
	;
	if base.Ui32(v108) < base.Ui32(int32(3)) {
		v169 = v122
		goto L43
	} else {
		goto L48
	}
L48:
	;
	if v117<<(uint(int32(2))%32)&int32(128) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return int32(0)
L50:
	;
	goto L51
L51:
	;
	if v108 == int32(3) {
		v169 = v122
		goto L43
	} else {
		goto L52
	}
L52:
	;
	if v117<<(uint(int32(3))%32)&int32(128) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	return int32(0)
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(v108) < base.Ui32(int32(5)) {
		v169 = v122
		goto L43
	} else {
		goto L56
	}
L56:
	;
	if v117<<(uint(int32(4))%32)&int32(128) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	return int32(0)
L58:
	;
	goto L59
L59:
	;
	if v108 == int32(5) {
		v169 = v122
		goto L43
	} else {
		goto L60
	}
L60:
	;
	if v117<<(uint(int32(5))%32)&int32(128) != 0 {
		v169 = int32(0)
		goto L43
	} else {
		goto L61
	}
L61:
	;
	if v108 != int32(7) {
		v169 = int32(1)
		goto L43
	} else {
		goto L62
	}
L62:
	;
	v169 = base.B2i32(v117&int32(2) == int32(0))
	goto L43
}
func F_network_subeq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v169 int32
	_ = v169
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = int32(1)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v15&v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = v13
	goto L6
L5:
	;
	v18 = int32(4)
	goto L6
L6:
	;
	v19 = v6 + v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v21 = int32(1)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v23&v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = v21
	goto L9
L8:
	;
	v26 = int32(4)
	goto L9
L9:
	;
	v27 = v11 + v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v20 != v28 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L12
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if base.Ui32(v33) < base.Ui32(v32) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	goto L15
L15:
	;
	v37 = int32(2)
	v38 = v19 + v37
	v40 = v27 + v37
	v42 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(4)) <= base.Ui32(v42) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if v104 != 0 {
		goto L34
	} else {
		goto L35
	}
L17:
	;
	v104 = int32(0)
	goto L16
L18:
	;
	v78 = v73
	v79 = v74
	v80 = v75
	goto L28
L19:
	;
	if (v38|v40)&int32(3) != 0 {
		v73 = v38
		v74 = v40
		v75 = v42
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v66 = v38
	v67 = v40
	v68 = v42
	goto L21
L21:
	;
	if v68 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L22:
	;
	v50 = v38
	v51 = v40
	v52 = v42
	goto L23
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v55 != v56 {
		v73 = v50
		v74 = v51
		v75 = v52
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v66 = v61
	v67 = v59
	v68 = v63
	goto L21
L25:
	;
	v58 = int32(4)
	v59 = v51 + v58
	v61 = v50 + v58
	v63 = v52 - v58
	if base.Ui32(int32(3)) < base.Ui32(v63) {
		v50 = v61
		v51 = v59
		v52 = v63
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v73 = v66
	v74 = v67
	v75 = v68
	goto L18
L28:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v83 == v84 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v104 = v83 - v84
	goto L16
L30:
	;
	v86 = int32(1)
	v91 = v80 - v86
	if v91 != 0 {
		v78 = v78 + v86
		v79 = v79 + v86
		v80 = v91
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L17
L34:
	;
	return int32(0)
L35:
	;
	goto L36
L36:
	;
	v108 = v32 & int32(7)
	if v108 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	return int32(1)
L38:
	;
	goto L39
L39:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v38))))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v40))))
	v117 = v114 ^ v116
	if base.Ui32(int32(127)) < base.Ui32(v117) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	return int32(0)
L41:
	;
	goto L42
L42:
	;
	v122 = int32(1)
	if v108 == v122 {
		v169 = v122
		goto L43
	} else {
		goto L44
	}
L43:
	;
	return v169
L44:
	;
	if v117<<(uint(int32(1))%32)&int32(128) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	return int32(0)
L46:
	;
	goto L47
L47:
	;
	if base.Ui32(v108) < base.Ui32(int32(3)) {
		v169 = v122
		goto L43
	} else {
		goto L48
	}
L48:
	;
	if v117<<(uint(int32(2))%32)&int32(128) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return int32(0)
L50:
	;
	goto L51
L51:
	;
	if v108 == int32(3) {
		v169 = v122
		goto L43
	} else {
		goto L52
	}
L52:
	;
	if v117<<(uint(int32(3))%32)&int32(128) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	return int32(0)
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(v108) < base.Ui32(int32(5)) {
		v169 = v122
		goto L43
	} else {
		goto L56
	}
L56:
	;
	if v117<<(uint(int32(4))%32)&int32(128) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	return int32(0)
L58:
	;
	goto L59
L59:
	;
	if v108 == int32(5) {
		v169 = v122
		goto L43
	} else {
		goto L60
	}
L60:
	;
	if v117<<(uint(int32(5))%32)&int32(128) != 0 {
		v169 = int32(0)
		goto L43
	} else {
		goto L61
	}
L61:
	;
	if v108 != int32(7) {
		v169 = int32(1)
		goto L43
	} else {
		goto L62
	}
L62:
	;
	v169 = base.B2i32(v117&int32(2) == int32(0))
	goto L43
}
func F_network_subset_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v8 != int32(461) {
		v47 = v2
		return v47
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if v11 == int32(0) {
			v47 = v2
			return v47
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			switch v14 - int32(15) {
			case 0, 2:
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				switch v23 - int32(927) {
				case 0:
					if v21 != 0 {
						v44 = v2
						v46 = v44
						v47 = v46
						return v47
					} else {
						v27 = F_match_network_subset(m, v19, v20, int32(0), v22)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							v46 = v27
							v47 = v46
							return v47
						}
					}
				case 1:
					if v21 != 0 {
						v44 = v2
						v46 = v44
						v47 = v46
						return v47
					} else {
						v32 = F_match_network_subset(m, v19, v20, int32(1), v22)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v46 = v32
							v47 = v46
							return v47
						}
					}
				case 2:
					if v21 != int32(1) {
						v44 = v2
						v46 = v44
						v47 = v46
						return v47
					} else {
						v37 = F_match_network_subset(m, v20, v19, int32(0), v22)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v46 = v37
							v47 = v46
							return v47
						}
					}
				case 3:
					if v21 != int32(1) {
						v44 = v2
						v46 = v44
						v47 = v46
						return v47
					} else {
						v42 = F_match_network_subset(m, v20, v19, int32(1), v22)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v44 = v42
							v46 = v44
							v47 = v46
							return v47
						}
					}
				default:
					v44 = v2
					v46 = v44
					v47 = v46
					return v47
				}
			default:
				v47 = v2
				return v47
			}
		}
	}
}
