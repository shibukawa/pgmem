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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int64
	_ = v127
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
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
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v285)
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
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218))))
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v219 != v220 {
		goto L47
	} else {
		goto L48
	}
L8:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v153 != v154 {
		v285 = v16
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
	if v33 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v56 = int32(-1) << (uint(int32(32)-v33) % 32)
	v57 = int32(16711935)
	v68 = base.I32_rotr(v56&v57, int32(8)) | base.I32_rotr(v56, int32(24))&v57
	goto L25
L24:
	;
	v68 = int32(0)
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
	goto L19
L26:
	;
	v74 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v74
	v77 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v77
	v86 = v74
	v89 = v33
	goto L27
L27:
	;
	v92 = v86 + (v25 + int32(8))
	v93 = int32(0)
	if v89 <= v93 {
		v103 = v93
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v125
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v127
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v129
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v131
	goto L19
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v92))) = uint8(v103)
	v105 = int32(0)
	v107 = v89 - int32(8)
	if v107 <= v105 {
		v117 = v105
		goto L32
	} else {
		goto L33
	}
L30:
	;
	if base.Ui32(int32(7)) < base.Ui32(v89) {
		v103 = int32(255)
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v103 = int32(255) << (uint(int32(8)-v89) % 32)
	goto L29
L32:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)) = uint8(v117)
	v119 = int32(16)
	v122 = v86 + int32(2)
	if v122 != v119 {
		v86 = v122
		v89 = v89 - v119
		goto L27
	} else {
		goto L35
	}
L33:
	;
	if base.Ui32(int32(7)) < base.Ui32(v107) {
		v117 = int32(255)
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v117 = int32(255) << (uint(int32(16)-v89) % 32)
	goto L32
L35:
	;
	goto L28
L36:
	;
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	switch v158 - int32(2) {
	case 0:
		goto L40
	default:
		goto L38
	case 8:
		goto L39
	}
L37:
	;
	if v214 == int32(0) {
		v285 = v16
		goto L4
	} else {
		goto L46
	}
L38:
	;
	v214 = int32(0)
	goto L37
L39:
	;
	v168 = int32(8)
	v169 = v8 + v168
	v171 = l0 + v168
	v173 = v152 + v168
	v175 = int32(0)
	goto L41
L40:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v214 = base.B2i32(v161&(v162^v163) == int32(0))
	goto L37
L41:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v169))))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v171))))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v173))))
	if v181&(v183^v185) != 0 {
		goto L38
	} else {
		goto L43
	}
L42:
	;
	v214 = int32(1)
	goto L37
L43:
	;
	v189 = v175 | int32(1)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171+v189))))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173+v189))))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v189))))
	if (v191^v193)&v196 != 0 {
		goto L38
	} else {
		goto L44
	}
L44:
	;
	v199 = v175 + int32(2)
	if v199 != int32(16) {
		v175 = v199
		goto L41
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	v285 = int32(1)
	goto L4
L47:
	;
	v285 = int32(0)
	goto L4
L48:
	;
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218))))
	switch v224 - int32(2) {
	case 0:
		goto L52
	default:
		goto L50
	case 8:
		goto L51
	}
L49:
	;
	if v280 == int32(0) {
		goto L47
	} else {
		goto L58
	}
L50:
	;
	v280 = int32(0)
	goto L49
L51:
	;
	v234 = int32(8)
	v235 = l1 + v234
	v237 = l0 + v234
	v239 = v218 + v234
	v241 = int32(0)
	goto L53
L52:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	v280 = base.B2i32(v227&(v228^v229) == int32(0))
	goto L49
L53:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241+v235))))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241+v237))))
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241+v239))))
	if v247&(v249^v251) != 0 {
		goto L50
	} else {
		goto L55
	}
L54:
	;
	v280 = int32(1)
	goto L49
L55:
	;
	v255 = v241 | int32(1)
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237+v255))))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239+v255))))
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235+v255))))
	if (v257^v259)&v262 != 0 {
		goto L50
	} else {
		goto L56
	}
L56:
	;
	v265 = v241 + int32(2)
	if v265 != int32(16) {
		v241 = v265
		goto L53
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	v285 = int32(1)
	goto L4
}
func F_network_cmp(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
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
					v41 = v23
				} else {
					v41 = v31
				}
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
				v44 = int32(base.Ui32(v42) >> (uint(int32(3)) % 32))
				v45 = F_memcmp(m, v35, v37, v44)
				mBase = m.M
				if v45 != 0 {
					v137 = v45
					v156 = v137
				} else {
					v47 = v42 & int32(7)
					if v47 == int32(0) {
						v128 = v38 - v39
						if v128 != 0 {
							v137 = v128
							v156 = v137
						} else {
							if v24 == int32(2) {
								v133 = int32(4)
							} else {
								v133 = int32(16)
							}
							v134 = F_memcmp(m, v35, v37, v133)
							mBase = m.M
							v156 = v134
						}
					} else {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v35))))
						v52 = int32(128)
						v53 = v51 & v52
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v37))))
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
								v128 = v38 - v39
								if v128 != 0 {
									v137 = v128
									v156 = v137
								} else {
									if v24 == int32(2) {
										v133 = int32(4)
									} else {
										v133 = int32(16)
									}
									v134 = F_memcmp(m, v35, v37, v133)
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
										v128 = v38 - v39
										if v128 != 0 {
											v137 = v128
											v156 = v137
										} else {
											if v24 == int32(2) {
												v133 = int32(4)
											} else {
												v133 = int32(16)
											}
											v134 = F_memcmp(m, v35, v37, v133)
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
												v128 = v38 - v39
												if v128 != 0 {
													v137 = v128
													v156 = v137
												} else {
													if v24 == int32(2) {
														v133 = int32(4)
													} else {
														v133 = int32(16)
													}
													v134 = F_memcmp(m, v35, v37, v133)
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
														v128 = v38 - v39
														if v128 != 0 {
															v137 = v128
															v156 = v137
														} else {
															if v24 == int32(2) {
																v133 = int32(4)
															} else {
																v133 = int32(16)
															}
															v134 = F_memcmp(m, v35, v37, v133)
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
																v128 = v38 - v39
																if v128 != 0 {
																	v137 = v128
																	v156 = v137
																} else {
																	if v24 == int32(2) {
																		v133 = int32(4)
																	} else {
																		v133 = int32(16)
																	}
																	v134 = F_memcmp(m, v35, v37, v133)
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
																		v128 = v38 - v39
																		if v128 != 0 {
																			v137 = v128
																			v156 = v137
																		} else {
																			if v24 == int32(2) {
																				v133 = int32(4)
																			} else {
																				v133 = int32(16)
																			}
																			v134 = F_memcmp(m, v35, v37, v133)
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
																			v128 = v38 - v39
																			if v128 != 0 {
																				v137 = v128
																				v156 = v137
																			} else {
																				if v24 == int32(2) {
																					v133 = int32(4)
																				} else {
																					v133 = int32(16)
																				}
																				v134 = F_memcmp(m, v35, v37, v133)
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
				v137 = v24 - v32
				v156 = v137
			}
			return base.I64_extend_i32_s(v156)
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	v10 = int32(1)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v12&v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if v261 != 0 {
		goto L71
	} else {
		goto L72
	}
L2:
	;
	return v253
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
	v253 = v17 - v25
	goto L2
L12:
	;
	v34 = v16
	goto L14
L13:
	;
	v34 = v24
	goto L14
L14:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	v37 = int32(base.Ui32(v35) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(4)) <= base.Ui32(v37) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	if v99 != 0 {
		v253 = v99
		goto L2
	} else {
		goto L33
	}
L16:
	;
	v99 = int32(0)
	goto L15
L17:
	;
	v73 = v68
	v74 = v69
	v75 = v70
	goto L27
L18:
	;
	if (v28|v30)&int32(3) != 0 {
		v68 = v28
		v69 = v30
		v70 = v37
		goto L17
	} else {
		goto L21
	}
L19:
	;
	v61 = v28
	v62 = v30
	v63 = v37
	goto L20
L20:
	;
	if v63 == int32(0) {
		goto L16
	} else {
		goto L26
	}
L21:
	;
	v45 = v28
	v46 = v30
	v47 = v37
	goto L22
L22:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v50 != v51 {
		v68 = v45
		v69 = v46
		v70 = v47
		goto L17
	} else {
		goto L24
	}
L23:
	;
	v61 = v56
	v62 = v54
	v63 = v58
	goto L20
L24:
	;
	v53 = int32(4)
	v54 = v46 + v53
	v56 = v45 + v53
	v58 = v47 - v53
	if base.Ui32(int32(3)) < base.Ui32(v58) {
		v45 = v56
		v46 = v54
		v47 = v58
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v68 = v61
	v69 = v62
	v70 = v63
	goto L17
L27:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v78 == v79 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v99 = v78 - v79
	goto L15
L29:
	;
	v81 = int32(1)
	v86 = v75 - v81
	if v86 != 0 {
		v73 = v73 + v81
		v74 = v74 + v81
		v75 = v86
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
	v101 = v35 & int32(7)
	if v101 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v182 = v31 - v32
	if v182 != 0 {
		v253 = v182
		goto L2
	} else {
		goto L49
	}
L35:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v28))))
	v106 = int32(128)
	v107 = v105 & v106
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v30))))
	if v107 != v109&v106 {
		v261 = v107
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v101 == int32(1) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v115 = int32(1)
	v117 = int32(128)
	v118 = v105 << (uint(v115) % 32) & v117
	if v118 != v109<<(uint(v115)%32)&v117 {
		v261 = v118
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if base.Ui32(v101) < base.Ui32(int32(3)) {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v126 = int32(2)
	v128 = int32(128)
	v129 = v105 << (uint(v126) % 32) & v128
	if v129 != v109<<(uint(v126)%32)&v128 {
		v261 = v129
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v101 == int32(3) {
		goto L34
	} else {
		goto L41
	}
L41:
	;
	v137 = int32(3)
	v139 = int32(128)
	v140 = v105 << (uint(v137) % 32) & v139
	if v140 != v109<<(uint(v137)%32)&v139 {
		v261 = v140
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if base.Ui32(v101) < base.Ui32(int32(5)) {
		goto L34
	} else {
		goto L43
	}
L43:
	;
	v148 = int32(4)
	v150 = int32(128)
	v151 = v105 << (uint(v148) % 32) & v150
	if v151 != v109<<(uint(v148)%32)&v150 {
		v261 = v151
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v101 == int32(5) {
		goto L34
	} else {
		goto L45
	}
L45:
	;
	v159 = int32(5)
	v161 = int32(128)
	v162 = v105 << (uint(v159) % 32) & v161
	if v162 != v109<<(uint(v159)%32)&v161 {
		v261 = v162
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v101 != int32(7) {
		goto L34
	} else {
		goto L47
	}
L47:
	;
	v170 = int32(6)
	v172 = int32(128)
	v173 = v105 << (uint(v170) % 32) & v172
	if v173 != v109<<(uint(v170)%32)&v172 {
		v261 = v173
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
	v187 = int32(4)
	goto L52
L51:
	;
	v187 = int32(16)
	goto L52
L52:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v187) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	return v249
L54:
	;
	v249 = int32(0)
	goto L53
L55:
	;
	v223 = v218
	v224 = v219
	v225 = v220
	goto L65
L56:
	;
	if (v28|v30)&int32(3) != 0 {
		v218 = v28
		v219 = v30
		v220 = v187
		goto L55
	} else {
		goto L59
	}
L57:
	;
	v211 = v28
	v212 = v30
	v213 = v187
	goto L58
L58:
	;
	if v213 == int32(0) {
		goto L54
	} else {
		goto L64
	}
L59:
	;
	v195 = v28
	v196 = v30
	v197 = v187
	goto L60
L60:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	if v200 != v201 {
		v218 = v195
		v219 = v196
		v220 = v197
		goto L55
	} else {
		goto L62
	}
L61:
	;
	v211 = v206
	v212 = v204
	v213 = v208
	goto L58
L62:
	;
	v203 = int32(4)
	v204 = v196 + v203
	v206 = v195 + v203
	v208 = v197 - v203
	if base.Ui32(int32(3)) < base.Ui32(v208) {
		v195 = v206
		v196 = v204
		v197 = v208
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v218 = v211
	v219 = v212
	v220 = v213
	goto L55
L65:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	if v228 == v229 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v249 = v228 - v229
	goto L53
L67:
	;
	v231 = int32(1)
	v236 = v225 - v231
	if v236 != 0 {
		v223 = v223 + v231
		v224 = v224 + v231
		v225 = v236
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
	v264 = int32(1)
	goto L73
L72:
	;
	v264 = int32(-1)
	goto L73
L73:
	;
	return v264
}
func F_network_eq(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
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
					v41 = v23
				} else {
					v41 = v31
				}
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
				v44 = int32(base.Ui32(v42) >> (uint(int32(3)) % 32))
				v45 = F_memcmp(m, v35, v37, v44)
				mBase = m.M
				if v45 != 0 {
					v137 = v45
					v156 = v137
				} else {
					v47 = v42 & int32(7)
					if v47 == int32(0) {
						v128 = v38 - v39
						if v128 != 0 {
							v137 = v128
							v156 = v137
						} else {
							if v24 == int32(2) {
								v133 = int32(4)
							} else {
								v133 = int32(16)
							}
							v134 = F_memcmp(m, v35, v37, v133)
							mBase = m.M
							v156 = v134
						}
					} else {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v35))))
						v52 = int32(128)
						v53 = v51 & v52
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v37))))
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
								v128 = v38 - v39
								if v128 != 0 {
									v137 = v128
									v156 = v137
								} else {
									if v24 == int32(2) {
										v133 = int32(4)
									} else {
										v133 = int32(16)
									}
									v134 = F_memcmp(m, v35, v37, v133)
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
										v128 = v38 - v39
										if v128 != 0 {
											v137 = v128
											v156 = v137
										} else {
											if v24 == int32(2) {
												v133 = int32(4)
											} else {
												v133 = int32(16)
											}
											v134 = F_memcmp(m, v35, v37, v133)
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
												v128 = v38 - v39
												if v128 != 0 {
													v137 = v128
													v156 = v137
												} else {
													if v24 == int32(2) {
														v133 = int32(4)
													} else {
														v133 = int32(16)
													}
													v134 = F_memcmp(m, v35, v37, v133)
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
														v128 = v38 - v39
														if v128 != 0 {
															v137 = v128
															v156 = v137
														} else {
															if v24 == int32(2) {
																v133 = int32(4)
															} else {
																v133 = int32(16)
															}
															v134 = F_memcmp(m, v35, v37, v133)
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
																v128 = v38 - v39
																if v128 != 0 {
																	v137 = v128
																	v156 = v137
																} else {
																	if v24 == int32(2) {
																		v133 = int32(4)
																	} else {
																		v133 = int32(16)
																	}
																	v134 = F_memcmp(m, v35, v37, v133)
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
																		v128 = v38 - v39
																		if v128 != 0 {
																			v137 = v128
																			v156 = v137
																		} else {
																			if v24 == int32(2) {
																				v133 = int32(4)
																			} else {
																				v133 = int32(16)
																			}
																			v134 = F_memcmp(m, v35, v37, v133)
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
																			v128 = v38 - v39
																			if v128 != 0 {
																				v137 = v128
																				v156 = v137
																			} else {
																				if v24 == int32(2) {
																					v133 = int32(4)
																				} else {
																					v133 = int32(16)
																				}
																				v134 = F_memcmp(m, v35, v37, v133)
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
				v137 = v24 - v32
				v156 = v137
			}
			return base.I64_extend_i32_u(base.B2i32(v156 == int32(0)))
		}
	}
}
func F_network_family(m *base.Module, l0 int32) int64 {
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
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
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
			v20 = int64(6)
		} else {
			v20 = int64(0)
		}
		if v17 == int32(2) {
			v23 = int64(4)
		} else {
			v23 = v20
		}
		return v23
	}
}
func F_network_ge(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
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
					v41 = v23
				} else {
					v41 = v31
				}
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
				v44 = int32(base.Ui32(v42) >> (uint(int32(3)) % 32))
				v45 = F_memcmp(m, v35, v37, v44)
				mBase = m.M
				if v45 != 0 {
					v137 = v45
					v156 = v137
				} else {
					v47 = v42 & int32(7)
					if v47 == int32(0) {
						v128 = v38 - v39
						if v128 != 0 {
							v137 = v128
							v156 = v137
						} else {
							if v24 == int32(2) {
								v133 = int32(4)
							} else {
								v133 = int32(16)
							}
							v134 = F_memcmp(m, v35, v37, v133)
							mBase = m.M
							v156 = v134
						}
					} else {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v35))))
						v52 = int32(128)
						v53 = v51 & v52
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v37))))
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
								v128 = v38 - v39
								if v128 != 0 {
									v137 = v128
									v156 = v137
								} else {
									if v24 == int32(2) {
										v133 = int32(4)
									} else {
										v133 = int32(16)
									}
									v134 = F_memcmp(m, v35, v37, v133)
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
										v128 = v38 - v39
										if v128 != 0 {
											v137 = v128
											v156 = v137
										} else {
											if v24 == int32(2) {
												v133 = int32(4)
											} else {
												v133 = int32(16)
											}
											v134 = F_memcmp(m, v35, v37, v133)
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
												v128 = v38 - v39
												if v128 != 0 {
													v137 = v128
													v156 = v137
												} else {
													if v24 == int32(2) {
														v133 = int32(4)
													} else {
														v133 = int32(16)
													}
													v134 = F_memcmp(m, v35, v37, v133)
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
														v128 = v38 - v39
														if v128 != 0 {
															v137 = v128
															v156 = v137
														} else {
															if v24 == int32(2) {
																v133 = int32(4)
															} else {
																v133 = int32(16)
															}
															v134 = F_memcmp(m, v35, v37, v133)
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
																v128 = v38 - v39
																if v128 != 0 {
																	v137 = v128
																	v156 = v137
																} else {
																	if v24 == int32(2) {
																		v133 = int32(4)
																	} else {
																		v133 = int32(16)
																	}
																	v134 = F_memcmp(m, v35, v37, v133)
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
																		v128 = v38 - v39
																		if v128 != 0 {
																			v137 = v128
																			v156 = v137
																		} else {
																			if v24 == int32(2) {
																				v133 = int32(4)
																			} else {
																				v133 = int32(16)
																			}
																			v134 = F_memcmp(m, v35, v37, v133)
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
																			v128 = v38 - v39
																			if v128 != 0 {
																				v137 = v128
																				v156 = v137
																			} else {
																				if v24 == int32(2) {
																					v133 = int32(4)
																				} else {
																					v133 = int32(16)
																				}
																				v134 = F_memcmp(m, v35, v37, v133)
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
				v137 = v24 - v32
				v156 = v137
			}
			return base.I64_extend_i32_u(base.B2i32(int32(0) <= v156))
		}
	}
}
func F_network_host(m *base.Module, l0 int32) int64 {
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
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
			return int64(0)
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
					return int64(0)
				} else {
					m.G0 = v6 - int32(-64)
					return base.I64_extend_i32_u(v40)
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50462850))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_network_host_0), int32(0))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_network_host_1), int32(1117), int32(_a_F_network_host_2))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
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
		}
	}
}
func F_network_network(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v13 = F_palloc0(m, int32(22))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			v16 = int32(1)
			v17 = v15 & v16
			v19 = int32(4)
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			v24 = v22 & v16
			if v24 != 0 {
				v25 = v16
			} else {
				v25 = v19
			}
			v26 = v8 + v25
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
			if v27 != 0 {
				if v17 != 0 {
					v30 = int32(1)
				} else {
					v30 = int32(4)
				}
				v32 = int32(2)
				v37 = v27
				v38 = int32(0)
				for {
					v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+(v26+v32)))))
					if base.Ui32(int32(7)) < base.Ui32(v37) {
						v53 = int32(-1)
					} else {
						v53 = int32(255) << (uint(int32(8)-v37) % 32)
					}
					v54 = v45 & v53
					*(*uint8)(unsafe.Add(mBase, uint32(v38+(v13+v30+v32)))) = uint8(v54)
					v58 = int32(8)
					if v37 <= v58 {
						v61 = v58
					} else {
						v61 = v37
					}
					v63 = v61 - int32(8)
					if v63 != 0 {
						v37 = v63
						v38 = v38 + int32(1)
						continue
					} else {
						break
					}
					break
				}
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
				v65 = int32(1)
				v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
				v74 = v64 & v65
				v76 = v67 & v65
			} else {
				v74 = v24
				v76 = v17
			}
			if v76 != 0 {
				v77 = v16
			} else {
				v77 = v19
			}
			v78 = v13 + v77
			if v74 != 0 {
				v81 = int32(1)
			} else {
				v81 = int32(4)
			}
			v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v81))))
			*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v83)
			v85 = int32(1)
			v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v87&v85 != 0 {
				v90 = v85
			} else {
				v90 = int32(4)
			}
			v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v90)+1)))
			*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)) = uint8(v92)
			if v83 == int32(2) {
				v98 = int32(40)
			} else {
				v98 = int32(88)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = v98
			return base.I64_extend_i32_u(v13)
		}
	}
}
func F_network_send(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
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
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	F_pq_begintypsend(m, v9)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = int32(1)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v17&v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = v15
	goto L5
L4:
	;
	v20 = int32(4)
	goto L5
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v20))))
	F_enlargeStringInfo(m, v9, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*uint8)(unsafe.Add(mBase, uint32(v26+v27))) = uint8(v22)
	v30 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v26 + v30
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v35&v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v38 = v30
	goto L9
L8:
	;
	v38 = int32(4)
	goto L9
L9:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v38)+1)))
	F_enlargeStringInfo(m, v9, int32(1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*uint8)(unsafe.Add(mBase, uint32(v44+v45))) = uint8(v40)
	v48 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v44 + v48
	F_enlargeStringInfo(m, v9, v48)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*uint8)(unsafe.Add(mBase, uint32(v54+v55))) = uint8(v2)
	v58 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v54 + v58
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v63&v58 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v66 = v58
	goto L14
L13:
	;
	v66 = int32(4)
	goto L14
L14:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v66))))
	F_enlargeStringInfo(m, v9, int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v68 == int32(2) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v79 = int32(4)
	goto L18
L17:
	;
	v79 = int32(16)
	goto L18
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v72+v73))) = uint8(v79)
	v81 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v72 + v81
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v86&v81 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v89 = v81
	goto L21
L20:
	;
	v89 = int32(4)
	goto L21
L21:
	;
	v94 = int32(0)
	goto L22
L22:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+(l0+v89+int32(2))))))
	F_enlargeStringInfo(m, v9, int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v117 << (uint(int32(2)) % 32)
	goto L26
L24:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*uint8)(unsafe.Add(mBase, uint32(v105+v106))) = uint8(v101)
	v109 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v105 + v109
	v113 = v94 + v109
	if v113 != v79 {
		v94 = v113
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	m.G0 = v9 + int32(16)
	return v116
}
func F_network_smaller(m *base.Module, l0 int32) int64 {
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
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v9 = F_pg_detoast_datum_packed(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
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
					v42 = v24
				} else {
					v42 = v32
				}
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
				v45 = int32(base.Ui32(v43) >> (uint(int32(3)) % 32))
				v46 = F_memcmp(m, v36, v38, v45)
				mBase = m.M
				if v46 != 0 {
					v138 = v46
					v157 = v138
				} else {
					v48 = v43 & int32(7)
					if v48 == int32(0) {
						v129 = v39 - v40
						if v129 != 0 {
							v138 = v129
							v157 = v138
						} else {
							if v25 == int32(2) {
								v134 = int32(4)
							} else {
								v134 = int32(16)
							}
							v135 = F_memcmp(m, v36, v38, v134)
							mBase = m.M
							v157 = v135
						}
					} else {
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v36))))
						v53 = int32(128)
						v54 = v52 & v53
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v38))))
						if v54 != v56&v53 {
							v145 = v54
							if v145 != 0 {
								v148 = int32(1)
							} else {
								v148 = int32(-1)
							}
							v157 = v148
						} else {
							if v48 == int32(1) {
								v129 = v39 - v40
								if v129 != 0 {
									v138 = v129
									v157 = v138
								} else {
									if v25 == int32(2) {
										v134 = int32(4)
									} else {
										v134 = int32(16)
									}
									v135 = F_memcmp(m, v36, v38, v134)
									mBase = m.M
									v157 = v135
								}
							} else {
								v62 = int32(1)
								v64 = int32(128)
								v65 = v52 << (uint(v62) % 32) & v64
								if v65 != v56<<(uint(v62)%32)&v64 {
									v145 = v65
									if v145 != 0 {
										v148 = int32(1)
									} else {
										v148 = int32(-1)
									}
									v157 = v148
								} else {
									if base.Ui32(v48) < base.Ui32(int32(3)) {
										v129 = v39 - v40
										if v129 != 0 {
											v138 = v129
											v157 = v138
										} else {
											if v25 == int32(2) {
												v134 = int32(4)
											} else {
												v134 = int32(16)
											}
											v135 = F_memcmp(m, v36, v38, v134)
											mBase = m.M
											v157 = v135
										}
									} else {
										v73 = int32(2)
										v75 = int32(128)
										v76 = v52 << (uint(v73) % 32) & v75
										if v76 != v56<<(uint(v73)%32)&v75 {
											v145 = v76
											if v145 != 0 {
												v148 = int32(1)
											} else {
												v148 = int32(-1)
											}
											v157 = v148
										} else {
											if v48 == int32(3) {
												v129 = v39 - v40
												if v129 != 0 {
													v138 = v129
													v157 = v138
												} else {
													if v25 == int32(2) {
														v134 = int32(4)
													} else {
														v134 = int32(16)
													}
													v135 = F_memcmp(m, v36, v38, v134)
													mBase = m.M
													v157 = v135
												}
											} else {
												v84 = int32(3)
												v86 = int32(128)
												v87 = v52 << (uint(v84) % 32) & v86
												if v87 != v56<<(uint(v84)%32)&v86 {
													v145 = v87
													if v145 != 0 {
														v148 = int32(1)
													} else {
														v148 = int32(-1)
													}
													v157 = v148
												} else {
													if base.Ui32(v48) < base.Ui32(int32(5)) {
														v129 = v39 - v40
														if v129 != 0 {
															v138 = v129
															v157 = v138
														} else {
															if v25 == int32(2) {
																v134 = int32(4)
															} else {
																v134 = int32(16)
															}
															v135 = F_memcmp(m, v36, v38, v134)
															mBase = m.M
															v157 = v135
														}
													} else {
														v95 = int32(4)
														v97 = int32(128)
														v98 = v52 << (uint(v95) % 32) & v97
														if v98 != v56<<(uint(v95)%32)&v97 {
															v145 = v98
															if v145 != 0 {
																v148 = int32(1)
															} else {
																v148 = int32(-1)
															}
															v157 = v148
														} else {
															if v48 == int32(5) {
																v129 = v39 - v40
																if v129 != 0 {
																	v138 = v129
																	v157 = v138
																} else {
																	if v25 == int32(2) {
																		v134 = int32(4)
																	} else {
																		v134 = int32(16)
																	}
																	v135 = F_memcmp(m, v36, v38, v134)
																	mBase = m.M
																	v157 = v135
																}
															} else {
																v106 = int32(5)
																v108 = int32(128)
																v109 = v52 << (uint(v106) % 32) & v108
																if v109 != v56<<(uint(v106)%32)&v108 {
																	v145 = v109
																	if v145 != 0 {
																		v148 = int32(1)
																	} else {
																		v148 = int32(-1)
																	}
																	v157 = v148
																} else {
																	if v48 != int32(7) {
																		v129 = v39 - v40
																		if v129 != 0 {
																			v138 = v129
																			v157 = v138
																		} else {
																			if v25 == int32(2) {
																				v134 = int32(4)
																			} else {
																				v134 = int32(16)
																			}
																			v135 = F_memcmp(m, v36, v38, v134)
																			mBase = m.M
																			v157 = v135
																		}
																	} else {
																		v117 = int32(6)
																		v119 = int32(128)
																		v120 = v52 << (uint(v117) % 32) & v119
																		if v120 != v56<<(uint(v117)%32)&v119 {
																			v145 = v120
																			if v145 != 0 {
																				v148 = int32(1)
																			} else {
																				v148 = int32(-1)
																			}
																			v157 = v148
																		} else {
																			v129 = v39 - v40
																			if v129 != 0 {
																				v138 = v129
																				v157 = v138
																			} else {
																				if v25 == int32(2) {
																					v134 = int32(4)
																				} else {
																					v134 = int32(16)
																				}
																				v135 = F_memcmp(m, v36, v38, v134)
																				mBase = m.M
																				v157 = v135
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
				v138 = v25 - v33
				v157 = v138
			}
			if v157 < int32(0) {
				v160 = v4
			} else {
				v160 = v9
			}
			return base.I64_extend_i32_u(v160)
		}
	}
}
func F_network_sub(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v155 int64
	_ = v155
	v5 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = int32(1)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v16&v14 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int64(1)
L5:
	;
	return v155
L6:
	;
	v19 = v14
	goto L8
L7:
	;
	v19 = int32(4)
	goto L8
L8:
	;
	v20 = v7 + v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v22 = int32(1)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v24&v22 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v27 = v22
	goto L11
L10:
	;
	v27 = int32(4)
	goto L11
L11:
	;
	v28 = v12 + v27
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v21 != v29 {
		v155 = v5
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if base.Ui32(v32) <= base.Ui32(v31) {
		v155 = v5
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v34 = int32(2)
	v35 = v20 + v34
	v37 = v28 + v34
	v39 = int32(base.Ui32(v31) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(4)) <= base.Ui32(v39) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v101 != 0 {
		v155 = v5
		goto L5
	} else {
		goto L32
	}
L15:
	;
	v101 = int32(0)
	goto L14
L16:
	;
	v75 = v70
	v76 = v71
	v77 = v72
	goto L26
L17:
	;
	if (v35|v37)&int32(3) != 0 {
		v70 = v35
		v71 = v37
		v72 = v39
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v63 = v35
	v64 = v37
	v65 = v39
	goto L19
L19:
	;
	if v65 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v47 = v35
	v48 = v37
	v49 = v39
	goto L21
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v52 != v53 {
		v70 = v47
		v71 = v48
		v72 = v49
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v63 = v58
	v64 = v56
	v65 = v60
	goto L19
L23:
	;
	v55 = int32(4)
	v56 = v48 + v55
	v58 = v47 + v55
	v60 = v49 - v55
	if base.Ui32(int32(3)) < base.Ui32(v60) {
		v47 = v58
		v48 = v56
		v49 = v60
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v70 = v63
	v71 = v64
	v72 = v65
	goto L16
L26:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v80 == v81 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v101 = v80 - v81
	goto L14
L28:
	;
	v83 = int32(1)
	v88 = v77 - v83
	if v88 != 0 {
		v75 = v75 + v83
		v76 = v76 + v83
		v77 = v88
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L15
L32:
	;
	v103 = v31 & int32(7)
	if v103 == int32(0) {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v39))))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v39))))
	v110 = v107 ^ v109
	if base.Ui32(int32(127)) < base.Ui32(v110) {
		v155 = v5
		goto L5
	} else {
		goto L34
	}
L34:
	;
	if v103 == int32(1) {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	if v110<<(uint(int32(1))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v103) < base.Ui32(int32(3)) {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	if v110<<(uint(int32(2))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L38
	}
L38:
	;
	if v103 == int32(3) {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	if v110<<(uint(int32(3))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L40
	}
L40:
	;
	if base.Ui32(v103) < base.Ui32(int32(5)) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v110<<(uint(int32(4))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L42
	}
L42:
	;
	if v103 == int32(5) {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	if v110<<(uint(int32(5))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L44
	}
L44:
	;
	if v103 != int32(7) {
		v155 = int64(1)
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v155 = base.I64_extend_i32_u(base.B2i32(v110&int32(2) == int32(0)))
	goto L5
}
func F_network_subeq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v155 int64
	_ = v155
	v5 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = int32(1)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v16&v14 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int64(1)
L5:
	;
	return v155
L6:
	;
	v19 = v14
	goto L8
L7:
	;
	v19 = int32(4)
	goto L8
L8:
	;
	v20 = v7 + v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v22 = int32(1)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v24&v22 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v27 = v22
	goto L11
L10:
	;
	v27 = int32(4)
	goto L11
L11:
	;
	v28 = v12 + v27
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v21 != v29 {
		v155 = v5
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if base.Ui32(v32) < base.Ui32(v31) {
		v155 = v5
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v34 = int32(2)
	v35 = v20 + v34
	v37 = v28 + v34
	v39 = int32(base.Ui32(v31) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(4)) <= base.Ui32(v39) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v101 != 0 {
		v155 = v5
		goto L5
	} else {
		goto L32
	}
L15:
	;
	v101 = int32(0)
	goto L14
L16:
	;
	v75 = v70
	v76 = v71
	v77 = v72
	goto L26
L17:
	;
	if (v35|v37)&int32(3) != 0 {
		v70 = v35
		v71 = v37
		v72 = v39
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v63 = v35
	v64 = v37
	v65 = v39
	goto L19
L19:
	;
	if v65 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v47 = v35
	v48 = v37
	v49 = v39
	goto L21
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v52 != v53 {
		v70 = v47
		v71 = v48
		v72 = v49
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v63 = v58
	v64 = v56
	v65 = v60
	goto L19
L23:
	;
	v55 = int32(4)
	v56 = v48 + v55
	v58 = v47 + v55
	v60 = v49 - v55
	if base.Ui32(int32(3)) < base.Ui32(v60) {
		v47 = v58
		v48 = v56
		v49 = v60
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v70 = v63
	v71 = v64
	v72 = v65
	goto L16
L26:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v80 == v81 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v101 = v80 - v81
	goto L14
L28:
	;
	v83 = int32(1)
	v88 = v77 - v83
	if v88 != 0 {
		v75 = v75 + v83
		v76 = v76 + v83
		v77 = v88
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L15
L32:
	;
	v103 = v31 & int32(7)
	if v103 == int32(0) {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v39))))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v39))))
	v110 = v107 ^ v109
	if base.Ui32(int32(127)) < base.Ui32(v110) {
		v155 = v5
		goto L5
	} else {
		goto L34
	}
L34:
	;
	if v103 == int32(1) {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	if v110<<(uint(int32(1))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v103) < base.Ui32(int32(3)) {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	if v110<<(uint(int32(2))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L38
	}
L38:
	;
	if v103 == int32(3) {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	if v110<<(uint(int32(3))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L40
	}
L40:
	;
	if base.Ui32(v103) < base.Ui32(int32(5)) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v110<<(uint(int32(4))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L42
	}
L42:
	;
	if v103 == int32(5) {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	if v110<<(uint(int32(5))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L44
	}
L44:
	;
	if v103 != int32(7) {
		v155 = int64(1)
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v155 = base.I64_extend_i32_u(base.B2i32(v110&int32(2) == int32(0)))
	goto L5
}
func F_network_subset_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v54 int64
	_ = v54
	v2 = int32(0)
	v7 = int64(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v9 != int32(469) {
		v54 = v7
		return v54
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		if v12 == int32(0) {
			v54 = v7
			return v54
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			switch v15 - int32(15) {
			case 0, 2:
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
				switch v24 - int32(927) {
				case 0:
					if v22 != 0 {
						v45 = v2
						v47 = v45
						v54 = base.I64_extend_i32_u(v47)
						return v54
					} else {
						v28 = F_match_network_subset(m, v20, v21, int32(0), v23)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int64(0)
						} else {
							v47 = v28
							v54 = base.I64_extend_i32_u(v47)
							return v54
						}
					}
				case 1:
					if v22 != 0 {
						v45 = v2
						v47 = v45
						v54 = base.I64_extend_i32_u(v47)
						return v54
					} else {
						v33 = F_match_network_subset(m, v20, v21, int32(1), v23)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							v47 = v33
							v54 = base.I64_extend_i32_u(v47)
							return v54
						}
					}
				case 2:
					if v22 != int32(1) {
						v45 = v2
						v47 = v45
						v54 = base.I64_extend_i32_u(v47)
						return v54
					} else {
						v38 = F_match_network_subset(m, v21, v20, int32(0), v23)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							v47 = v38
							v54 = base.I64_extend_i32_u(v47)
							return v54
						}
					}
				case 3:
					if v22 != int32(1) {
						v45 = v2
						v47 = v45
						v54 = base.I64_extend_i32_u(v47)
						return v54
					} else {
						v43 = F_match_network_subset(m, v21, v20, int32(1), v23)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							v45 = v43
							v47 = v45
							v54 = base.I64_extend_i32_u(v47)
							return v54
						}
					}
				default:
					v45 = v2
					v47 = v45
					v54 = base.I64_extend_i32_u(v47)
					return v54
				}
			default:
				v54 = v7
				return v54
			}
		}
	}
}
