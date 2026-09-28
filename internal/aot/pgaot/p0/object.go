package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetNewObjectId(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetNewObjectId[0])))
	if v5 == int32(1) {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewObjectId[1]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+308))
		v13 = base.B2i32(v11 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_GetNewObjectId[0])) = uint8(v13)
		v15 = v13
	} else {
		v15 = int32(0)
	}
	if v15 == int32(0) {
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewObjectId[2]))
		v23 = F_LWLockAcquire(m, v19+int32(256), int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewObjectId[3]))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
			if base.Ui32(int32(_a_F_GetNewObjectId_0)) < base.Ui32(v29) {
				v45 = v28
			} else {
				v33 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetNewObjectId[4])))
				if base.B2i32(v33 == int32(0))&base.B2i32(base.Ui32(int32(_a_F_GetNewObjectId_1)) < base.Ui32(v29)) != 0 {
					v45 = v28
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(_a_F_GetNewObjectId_2)
					v42 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewObjectId[3]))
					*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = int32(0)
					v45 = v42
				}
			}
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
			if v46 == int32(0) {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
				v52 = m.G0
				v54 = v52 - int32(16)
				m.G0 = v54
				*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v49 - int32(-8192)
				F_XLogBeginInsert(m)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					F_XLogRegisterData(m, v54+int32(12), int32(4))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v66 = F_XLogInsert(m, int32(0), int32(48))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							m.G0 = v54 + int32(16)
							v72 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewObjectId[3]))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = int32(_a_F_GetNewObjectId_3)
							v75 = v72
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
							v78 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v75))) = v77 + v78
							v82 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewObjectId[3]))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v83 - v78
							v88 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewObjectId[2]))
							F_LWLockRelease(m, v88+int32(256))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								return v77
							}
						}
					}
				}
			} else {
				v75 = v45
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
				v78 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v75))) = v77 + v78
				v82 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewObjectId[3]))
				v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v83 - v78
				v88 = *(*int32)(unsafe.Add(mBase, _c_F_GetNewObjectId[2]))
				F_LWLockRelease(m, v88+int32(256))
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int32(0)
				} else {
					return v77
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v97 = m.ExcPending
		if v97 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_GetNewObjectId_4), int32(0))
			mBase = m.M
			v101 = m.ExcPending
			if v101 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_GetNewObjectId_5), int32(560), int32(_a_F_GetNewObjectId_6))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
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
func F_check_object_ownership(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	switch l1 {
	case 0, 27, 48, 49:
		v208 = F_superuser_arg(m, l0)
		mBase = m.M
		v209 = m.ExcPending
		if v209 != 0 {
			return
		} else {
			if v208 != 0 {
				m.G0 = v9 + int32(112)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v213 = m.ExcPending
				if v213 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v216 = m.ExcPending
					if v216 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_check_object_ownership_0), int32(0))
						mBase = m.M
						v220 = m.ExcPending
						if v220 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_object_ownership_1), int32(2600), int32(_a_F_check_object_ownership_2))
							mBase = m.M
							v225 = m.ExcPending
							if v225 != 0 {
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
	case 1, 19, 25, 29, 35:
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v49 = F_object_ownercheck(m, v47, v48, l0)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return
		} else {
			if v49 != 0 {
				m.G0 = v9 + int32(112)
				return
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
				v53 = F_NameListToString(m, v52)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					F_aclcheck_error(m, int32(2), l1, v53)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						m.G0 = v9 + int32(112)
						return
					}
				}
			}
		}
	case 2, 3, 10, 11, 31, 32, 33, 51:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v229 = m.ExcPending
		if v229 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = l1
			F_errmsg_internal(m, int32(_a_F_check_object_ownership_3), v9+int32(96))
			mBase = m.M
			v235 = m.ExcPending
			if v235 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_check_object_ownership_1), int32(2611), int32(_a_F_check_object_ownership_2))
				mBase = m.M
				v240 = m.ExcPending
				if v240 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 4, 12, 50:
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v23 = F_object_ownercheck(m, v21, v22, l0)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			if v23 != 0 {
				m.G0 = v9 + int32(112)
				return
			} else {
				F_aclcheck_error_type(m, int32(2), v22)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					m.G0 = v9 + int32(112)
					return
				}
			}
		}
	case 5:
		v98 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
		v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
		v101 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
		v102 = F_typenameTypeId(m, int32(0), v101)
		mBase = m.M
		v103 = m.ExcPending
		if v103 != 0 {
			return
		} else {
			v105 = F_typenameTypeId(m, int32(0), v99)
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return
			} else {
				v108 = F_object_ownercheck(m, int32(1247), v102, l0)
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
					return
				} else {
					if v108 != 0 {
						m.G0 = v9 + int32(112)
						return
					} else {
						v111 = F_object_ownercheck(m, int32(1247), v105, l0)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return
						} else {
							if v111 != 0 {
								m.G0 = v9 + int32(112)
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return
								} else {
									F_errcode(m, int32(16797828))
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return
									} else {
										v120 = F_format_type_be(m, v102)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return
										} else {
											v122 = F_format_type_be(m, v105)
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v122
												*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v120
												F_errmsg(m, int32(_a_F_check_object_ownership_4), v9+int32(32))
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_check_object_ownership_1), int32(2545), int32(_a_F_check_object_ownership_2))
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
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
					}
				}
			}
		}
	case 6, 18, 20, 23, 28, 36, 38, 41, 42, 45, 52:
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
		v13 = F_object_ownercheck(m, int32(1259), v12, l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			if v13 != 0 {
				m.G0 = v9 + int32(112)
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l4)+48))
				F_aclcheck_error(m, int32(2), l1, v16+int32(4))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					m.G0 = v9 + int32(112)
					return
				}
			}
		}
	case 7, 8, 24, 26, 40, 46, 47:
		v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v67 = F_object_ownercheck(m, v65, v66, l0)
		mBase = m.M
		v68 = m.ExcPending
		if v68 != 0 {
			return
		} else {
			if v67 != 0 {
				m.G0 = v9 + int32(112)
				return
			} else {
				v70 = F_NameListToString(m, l3)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					F_aclcheck_error(m, int32(2), l1, v70)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						m.G0 = v9 + int32(112)
						return
					}
				}
			}
		}
	case 9, 14, 15, 16, 17, 21, 30, 37, 39, 43:
		v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v59 = F_object_ownercheck(m, v57, v58, l0)
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return
		} else {
			if v59 != 0 {
				m.G0 = v9 + int32(112)
				return
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
				F_aclcheck_error(m, int32(2), l1, v62)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					m.G0 = v9 + int32(112)
					return
				}
			}
		}
	case 13:
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v31 = F_SearchSysCache1(m, int32(19), base.I64_extend_i32_u(v29))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			if v31 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v244 = m.ExcPending
				if v244 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v29
					F_errmsg_internal(m, int32(_a_F_check_object_ownership_5), v9)
					mBase = m.M
					v248 = m.ExcPending
					if v248 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_object_ownership_1), int32(2475), int32(_a_F_check_object_ownership_2))
						mBase = m.M
						v253 = m.ExcPending
						if v253 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v35+v36)+84))
				F_ReleaseCatCache(m, v31)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					v42 = F_object_ownercheck(m, int32(1247), v38, l0)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						if v42 != 0 {
							m.G0 = v9 + int32(112)
							return
						} else {
							F_aclcheck_error_type(m, int32(2), v38)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								m.G0 = v9 + int32(112)
								return
							}
						}
					}
				}
			}
		}
	case 22:
		v75 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_object_ownership[0])))
		if v75 != 0 {
			m.G0 = v9 + int32(112)
			return
		} else {
			v76 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v78 = F_object_ownercheck(m, v76, v77, l0)
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return
			} else {
				if v78 != 0 {
					m.G0 = v9 + int32(112)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						F_errcode(m, int32(16797828))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v77
							F_errmsg(m, int32(_a_F_check_object_ownership_6), v9+int32(16))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_check_object_ownership_1), int32(2529), int32(_a_F_check_object_ownership_2))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
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
	case 34:
		v147 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v148 = F_superuser_arg(m, v147)
		mBase = m.M
		v149 = m.ExcPending
		if v149 != 0 {
			return
		} else {
			if v148 != 0 {
				v150 = F_superuser_arg(m, l0)
				mBase = m.M
				v151 = m.ExcPending
				if v151 != 0 {
					return
				} else {
					if v150 != 0 {
						m.G0 = v9 + int32(112)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v155 = m.ExcPending
						if v155 != 0 {
							return
						} else {
							F_errcode(m, int32(16797828))
							mBase = m.M
							v158 = m.ExcPending
							if v158 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_check_object_ownership_7), int32(0))
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = int32(_a_F_check_object_ownership_8)
									v168 = F_errdetail(m, int32(_a_F_check_object_ownership_9), v9+int32(48))
									mBase = m.M
									v169 = m.ExcPending
									if v169 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_check_object_ownership_1), int32(2572), int32(_a_F_check_object_ownership_2))
										mBase = m.M
										v174 = m.ExcPending
										if v174 != 0 {
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
			} else {
				v175 = F_has_createrole_privilege(m, l0)
				mBase = m.M
				v176 = m.ExcPending
				if v176 != 0 {
					return
				} else {
					if v175 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v257 = m.ExcPending
						if v257 != 0 {
							return
						} else {
							F_errcode(m, int32(16797828))
							mBase = m.M
							v260 = m.ExcPending
							if v260 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_check_object_ownership_7), int32(0))
								mBase = m.M
								v264 = m.ExcPending
								if v264 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = int32(_a_F_check_object_ownership_10)
									v270 = F_errdetail(m, int32(_a_F_check_object_ownership_9), v9+int32(80))
									mBase = m.M
									v271 = m.ExcPending
									if v271 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_check_object_ownership_1), int32(2581), int32(_a_F_check_object_ownership_2))
										mBase = m.M
										v276 = m.ExcPending
										if v276 != 0 {
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
					} else {
						v179 = F_is_admin_of_role(m, l0, v147)
						mBase = m.M
						v180 = m.ExcPending
						if v180 != 0 {
							return
						} else {
							if v179 != 0 {
								m.G0 = v9 + int32(112)
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return
								} else {
									F_errcode(m, int32(16797828))
									mBase = m.M
									v187 = m.ExcPending
									if v187 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_check_object_ownership_7), int32(0))
										mBase = m.M
										v191 = m.ExcPending
										if v191 != 0 {
											return
										} else {
											v193 = F_GetUserNameFromId(m, v147, int32(1))
											mBase = m.M
											v194 = m.ExcPending
											if v194 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v193
												*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = int32(_a_F_check_object_ownership_11)
												v201 = F_errdetail(m, int32(_a_F_check_object_ownership_12), v9-int32(-64))
												mBase = m.M
												v202 = m.ExcPending
												if v202 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_check_object_ownership_1), int32(2589), int32(_a_F_check_object_ownership_2))
													mBase = m.M
													v207 = m.ExcPending
													if v207 != 0 {
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
					}
				}
			}
		}
	case 44:
		v138 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
		v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
		v140 = F_typenameTypeId(m, int32(0), v139)
		mBase = m.M
		v141 = m.ExcPending
		if v141 != 0 {
			return
		} else {
			v142 = F_object_ownercheck(m, int32(1247), v140, l0)
			mBase = m.M
			v143 = m.ExcPending
			if v143 != 0 {
				return
			} else {
				if v142 != 0 {
					m.G0 = v9 + int32(112)
					return
				} else {
					F_aclcheck_error_type(m, int32(2), v140)
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return
					} else {
						m.G0 = v9 + int32(112)
						return
					}
				}
			}
		}
	default:
		m.G0 = v9 + int32(112)
		return
	}
}
func F_get_object_address_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(1247)
	v16 = F_LookupTypeName(m, v5, l2, l3)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		if v16 == int32(0) {
			if l3 != 0 {
				m.G0 = v9 + int32(32)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v27 = F_TypeNameToString(m, l2)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v27
							F_errmsg(m, int32(_a_F_get_object_address_type_0), v9)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_get_object_address_type_1), int32(1633), int32(_a_F_get_object_address_type_2))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
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
		} else {
			v38 = F_typeTypeId(m, v16)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v38
				if l1 == int32(12) {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
					v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v44)+79)))
					if v46 != int32(100) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							F_errcode(m, int32(151027844))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								v63 = F_TypeNameToString(m, l2)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v63
									F_errmsg(m, int32(_a_F_get_object_address_type_3), v9+int32(16))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_get_object_address_type_1), int32(1644), int32(_a_F_get_object_address_type_2))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
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
					} else {
						F_ReleaseCatCache(m, v16)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							m.G0 = v9 + int32(32)
							return
						}
					}
				} else {
					F_ReleaseCatCache(m, v16)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				}
			}
		}
	}
}
func F_get_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v11 < v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = int32(1)
	v15 = v10 - v14
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v15))))
	if v17 != v14 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v20 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20+v15<<(uint(int32(2))%32))))
	if v26 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if base.B2i32(v31 == int32(0))|base.B2i32(v31 != v34) != 0 {
		v52 = v31
		v53 = v34
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v52-v53 != 0 {
		goto L1
	} else {
		goto L13
	}
L7:
	;
	goto L6
L8:
	;
	v37 = l1
	v38 = v26
	goto L9
L9:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v42 == int32(0) {
		v52 = v42
		v53 = v41
		goto L7
	} else {
		goto L11
	}
L10:
	;
	v52 = v42
	v53 = v41
	goto L7
L11:
	;
	v45 = int32(1)
	if v42 == v41 {
		v37 = v37 + v45
		v38 = v38 + v45
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	if v10 < v11 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v57 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10+v13))) = uint8(v57)
	return int32(0)
L15:
	;
	goto L16
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v63 != int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v73
	goto L1
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if v66 != int32(1) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v69 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v69)
	return int32(0)
}
