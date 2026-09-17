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
								F_appendStringInfoString(m, v42, int32(_a_F_ExplainOpenGroup_0))
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
							F_appendStringInfoString(m, v42, int32(_a_F_ExplainOpenGroup_0))
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
				F_appendStringInfo(m, v78, int32(_a_F_ExplainOpenGroup_1), v8)
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
				F_appendStringInfoString(m, v78, int32(_a_F_ExplainOpenGroup_2))
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
						F_appendStringInfo(m, v78, int32(_a_F_ExplainOpenGroup_1), v8)
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
						F_appendStringInfoString(m, v78, int32(_a_F_ExplainOpenGroup_2))
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
	var v7 int32
	_ = v7
	Fn13827(m, l0, l1, l2, l3, int32(_a_F_ExplainPropertyInteger_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
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
		goto L52
	} else {
		goto L53
	}
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v74 = int32(_a_F_ExplainXMLTag_0)
	v75 = base.I32_extend8_s(v70)
	v76 = int32(66)
	goto L25
L21:
	;
	goto L19
L22:
	;
	if v181 != 0 {
		goto L47
	} else {
		goto L48
	}
L23:
	;
	v181 = int32(0)
	goto L22
L24:
	;
	v159 = v152
	v161 = v154
	goto L41
L25:
	;
	goto L32
L32:
	;
	v115 = v75 & int32(255)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExplainXMLTag[0])))
	if base.B2i32(v115 == v116)|int32(0) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v125 = v74
	v127 = v76
	goto L36
L34:
	;
	v145 = v74
	v147 = v76
	goto L35
L35:
	;
	if v147 == int32(0) {
		goto L23
	} else {
		goto L40
	}
L36:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v132 = v131 ^ v115*int32(16843009)
	v135 = int32(-2139062144)
	if (int32(16843008)-v132|v132)&v135 != v135 {
		v152 = v125
		v154 = v127
		goto L24
	} else {
		goto L38
	}
L37:
	;
	v145 = v140
	v147 = v142
	goto L35
L38:
	;
	v139 = int32(4)
	v140 = v125 + v139
	v142 = v127 - v139
	if base.Ui32(int32(3)) < base.Ui32(v142) {
		v125 = v140
		v127 = v142
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v152 = v145
	v154 = v147
	goto L24
L41:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v75&int32(255) == v164 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L23
L43:
	;
	v181 = v159
	goto L22
L44:
	;
	goto L45
L45:
	;
	v166 = int32(1)
	v169 = v161 - v166
	if v169 != 0 {
		v159 = v159 + v166
		v161 = v169
		goto L41
	} else {
		goto L46
	}
L46:
	;
	goto L42
L47:
	;
	v182 = v70
	goto L49
L48:
	;
	v182 = int32(45)
	goto L49
L49:
	;
	F_appendStringInfoChar(m, v72, base.I32_extend8_s(v182))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	if v186 != 0 {
		v67 = v67 + int32(1)
		v70 = v186
		goto L20
	} else {
		goto L51
	}
L51:
	;
	goto L21
L52:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v196, int32(_a_F_ExplainXMLTag_1))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L4
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	if v204 <= v201+int32(1) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L54
L56:
	;
	if base.Ui32(l1) <= base.Ui32(int32(3)) {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	F_appendStringInfoChar(m, v200, int32(62))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	v211 = int32(62)
	*(*uint8)(unsafe.Add(mBase, uint32(v209+v201))) = uint8(v211)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v216 = v214 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v213)+4)) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v220 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v218+v216))) = uint8(v220)
	goto L56
L60:
	;
	goto L56
L61:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v226)+8))
	if v230 <= v227+int32(1) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	return
L64:
	;
	F_appendStringInfoChar(m, v226, int32(10))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v237 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v235+v227))) = uint8(v237)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	v242 = v240 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v239)+4)) = v242
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	v246 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v244+v242))) = uint8(v246)
	goto L63
L67:
	;
	return
}
