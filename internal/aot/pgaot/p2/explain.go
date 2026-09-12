package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExplainCloseGroup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	switch v4 - int32(1) {
	case 0:
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
		v8 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v7 - v8
		F_ExplainXMLTag(m, l0, v8, l2)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			return
		}
	case 1:
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
		*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v14 - int32(1)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		F_appendStringInfoChar(m, v18, int32(10))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
			F_appendStringInfoSpaces(m, v22, v23<<(uint(int32(1))%32))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				if l1 != 0 {
					v31 = int32(125)
				} else {
					v31 = int32(93)
				}
				F_appendStringInfoChar(m, v28, v31)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
					v35 = F_list_delete_first(m, v34)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v35
						return
					}
				}
			}
		}
	case 2:
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
		*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v38 - int32(1)
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
		v43 = F_list_delete_first(m, v42)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v43
			return
		}
	default:
		return
	}
}
func F_ExplainOpenGroup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	switch v10 - int32(1) {
	case 0:
		F_ExplainXMLTag(m, l0, int32(0), l3)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v16 + int32(1)
			m.G0 = v8 + int32(16)
			return
		}
	case 1:
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		if v22 != 0 {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			F_appendStringInfoChar(m, v23, int32(44))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				F_appendStringInfoChar(m, v29, int32(10))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
					F_appendStringInfoSpaces(m, v33, v34<<(uint(int32(1))%32))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						if l1 != 0 {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							F_escape_json(m, v39, l1)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								F_appendStringInfoString(m, v42, int32(705584))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									if l2 != 0 {
										v49 = int32(123)
									} else {
										v49 = int32(91)
									}
									F_appendStringInfoChar(m, v46, v49)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
										v54 = F_lcons_int(m, int32(0), v53)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l3)+28)) = v54
											v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
											*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v57 + int32(1)
											m.G0 = v8 + int32(16)
											return
										}
									}
								}
							}
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							if l2 != 0 {
								v49 = int32(123)
							} else {
								v49 = int32(91)
							}
							F_appendStringInfoChar(m, v46, v49)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
								v54 = F_lcons_int(m, int32(0), v53)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l3)+28)) = v54
									v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
									*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v57 + int32(1)
									m.G0 = v8 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(1)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			F_appendStringInfoChar(m, v29, int32(10))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
				F_appendStringInfoSpaces(m, v33, v34<<(uint(int32(1))%32))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					if l1 != 0 {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						F_escape_json(m, v39, l1)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							F_appendStringInfoString(m, v42, int32(705584))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								if l2 != 0 {
									v49 = int32(123)
								} else {
									v49 = int32(91)
								}
								F_appendStringInfoChar(m, v46, v49)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
									v54 = F_lcons_int(m, int32(0), v53)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l3)+28)) = v54
										v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v57 + int32(1)
										m.G0 = v8 + int32(16)
										return
									}
								}
							}
						}
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						if l2 != 0 {
							v49 = int32(123)
						} else {
							v49 = int32(91)
						}
						F_appendStringInfoChar(m, v46, v49)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
							v54 = F_lcons_int(m, int32(0), v53)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3)+28)) = v54
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
								*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v57 + int32(1)
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	case 2:
		v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
		if v63 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v62))) = int32(1)
			v78 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			if l1 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
				F_appendStringInfo(m, v78, int32(705227), v8)
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return
				} else {
					v88 = int32(1)
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
					v90 = F_lcons_int(m, v88, v89)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3)+28)) = v90
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
						*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v93 + int32(1)
						m.G0 = v8 + int32(16)
						return
					}
				}
			} else {
				F_appendStringInfoString(m, v78, int32(705593))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					v88 = int32(0)
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
					v90 = F_lcons_int(m, v88, v89)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3)+28)) = v90
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
						*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v93 + int32(1)
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		} else {
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			F_appendStringInfoChar(m, v68, int32(10))
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return
			} else {
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				v73 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
				F_appendStringInfoSpaces(m, v72, v73<<(uint(int32(1))%32))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					if l1 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
						F_appendStringInfo(m, v78, int32(705227), v8)
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return
						} else {
							v88 = int32(1)
							v89 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
							v90 = F_lcons_int(m, v88, v89)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3)+28)) = v90
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
								*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v93 + int32(1)
								m.G0 = v8 + int32(16)
								return
							}
						}
					} else {
						F_appendStringInfoString(m, v78, int32(705593))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return
						} else {
							v88 = int32(0)
							v89 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
							v90 = F_lcons_int(m, v88, v89)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3)+28)) = v90
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
								*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v93 + int32(1)
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	default:
		m.G0 = v8 + int32(16)
		return
	}
}
func F_ExplainPropertyInteger(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = l2
	v15 = F_pg_snprintf(m, v8+int32(16), int32(32), int32(411408), v8)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		F_ExplainProperty(m, l0, l1, v8+int32(16), int32(1), l3)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			m.G0 = v8 + int32(48)
			return
		}
	}
}
func F_ExplainXMLTag(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	if base.Ui32(l1) <= base.Ui32(int32(3)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	F_appendStringInfoSpaces(m, v8, v9<<(uint(int32(1))%32))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v18 <= v15+int32(1) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	if l1&int32(1) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	F_appendStringInfoChar(m, v14, int32(60))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v25 = int32(60)
	*(*uint8)(unsafe.Add(mBase, uint32(v23+v15))) = uint8(v25)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v30 = v28 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v34 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32+v30))) = uint8(v34)
	goto L6
L10:
	;
	goto L6
L11:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v66 != 0 {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v46 <= v43+int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_appendStringInfoChar(m, v42, int32(47))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v53 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v51+v43))) = uint8(v53)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v58 = v56 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v62 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60+v58))) = uint8(v62)
	goto L11
L16:
	;
	goto L11
L17:
	;
	v67 = l0
	v70 = v66
	goto L20
L18:
	;
	goto L19
L19:
	;
	if l1&int32(2) != 0 {
		goto L53
	} else {
		goto L54
	}
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v75 = base.I32_extend8_s(v70)
	goto L26
L21:
	;
	goto L19
L22:
	;
	if v179 != 0 {
		goto L48
	} else {
		goto L49
	}
L23:
	;
	v179 = int32(0)
	goto L22
L24:
	;
	v157 = v150
	v159 = v152
	goto L42
L25:
	;
	if base.B2i32(v97 != v98) == int32(0) {
		goto L23
	} else {
		goto L33
	}
L26:
	;
	goto L27
L27:
	;
	v89 = int32(615971)
	v91 = int32(66)
	goto L28
L28:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v94 == v75&int32(255) {
		v150 = v89
		v152 = v91
		goto L24
	} else {
		goto L30
	}
L29:
	;
	goto L25
L30:
	;
	v96 = int32(1)
	v97 = v91 - v96
	v98 = int32(0)
	v101 = v89 + v96
	if v101&int32(3) == v98 {
		goto L25
	} else {
		goto L31
	}
L31:
	;
	if v97 != 0 {
		v89 = v101
		v91 = v97
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v113 == v75&int32(255) {
		v143 = v101
		v145 = v97
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v145 == int32(0) {
		goto L23
	} else {
		goto L41
	}
L35:
	;
	if base.Ui32(v97) < base.Ui32(int32(4)) {
		v143 = v101
		v145 = v97
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v123 = v101
	v125 = v97
	goto L37
L37:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v130 = v129 ^ v75&int32(255)*int32(16843009)
	v133 = int32(-2139062144)
	if (int32(16843008)-v130|v130)&v133 != v133 {
		v150 = v123
		v152 = v125
		goto L24
	} else {
		goto L39
	}
L38:
	;
	v143 = v138
	v145 = v140
	goto L34
L39:
	;
	v137 = int32(4)
	v138 = v123 + v137
	v140 = v125 - v137
	if base.Ui32(int32(3)) < base.Ui32(v140) {
		v123 = v138
		v125 = v140
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v150 = v143
	v152 = v145
	goto L24
L42:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v75&int32(255) == v162 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L23
L44:
	;
	v179 = v157
	goto L22
L45:
	;
	goto L46
L46:
	;
	v164 = int32(1)
	v167 = v159 - v164
	if v167 != 0 {
		v157 = v157 + v164
		v159 = v167
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	v180 = v70
	goto L50
L49:
	;
	v180 = int32(45)
	goto L50
L50:
	;
	F_appendStringInfoChar(m, v72, base.I32_extend8_s(v180))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v185 = v67 + int32(1)
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if v186 != 0 {
		v67 = v185
		v70 = v186
		goto L20
	} else {
		goto L52
	}
L52:
	;
	goto L21
L53:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v194, int32(532354))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198)+8))
	if v202 <= v199+int32(1) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L55
L57:
	;
	if base.Ui32(l1) <= base.Ui32(int32(3)) {
		goto L62
	} else {
		goto L63
	}
L58:
	;
	F_appendStringInfoChar(m, v198, int32(62))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v209 = int32(62)
	*(*uint8)(unsafe.Add(mBase, uint32(v207+v199))) = uint8(v209)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v214 = v212 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v211)+4)) = v214
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v218 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v216+v214))) = uint8(v218)
	goto L57
L61:
	;
	goto L57
L62:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v224)+8))
	if v228 <= v225+int32(1) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	return
L65:
	;
	F_appendStringInfoChar(m, v224, int32(10))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v235 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v233+v225))) = uint8(v235)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	v240 = v238 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v240
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v244 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v242+v240))) = uint8(v244)
	goto L64
L68:
	;
	return
}
