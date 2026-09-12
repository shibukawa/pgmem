package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetPublicationRelations(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v13 = F_table_open(m, int32(6106), int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_ScanKeyInit(m, v9, int32(2), int32(3), int32(184), l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = int32(0)
	v24 = int32(1)
	v27 = F_systable_beginscan(m, v13, int32(6116), v24, v22, v24, v9)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = v22
	goto L5
L5:
	;
	v35 = F_systable_getnext(m, v27)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	F_systable_endscan(m, v27)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	if v35 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37+v38)+8))
	v41 = F_GetPubPartitionOptionRelations(m, v29, l1, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	goto L6
L11:
	;
	v29 = v41
	goto L5
L12:
	;
	F_sequence_close(m, v13, int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_list_sort(m, v29, int32(467))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v29 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	m.G0 = v9 + int32(48)
	return v29
L16:
	;
	goto L15
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v62 < int32(2) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v65 = int32(1)
	v67 = v62 - v65
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v62 == int32(2) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v67&v65 == int32(0) {
		v152 = v126
		goto L32
	} else {
		goto L33
	}
L20:
	;
	v126 = int32(0)
	v128 = v65
	goto L19
L21:
	;
	goto L22
L22:
	;
	v78 = int32(0)
	v81 = v78
	v83 = v65
	v84 = v78
	goto L23
L23:
	;
	v90 = int32(2)
	v91 = v83 << (uint(v90) % 32)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v70+v91)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v70+v81<<(uint(v90)%32))))
	if v93 != v97 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v126 = v119
	v128 = v121
	goto L19
L25:
	;
	v100 = v81 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v70+v100<<(uint(int32(2))%32)))) = v93
	v105 = v100
	goto L27
L26:
	;
	v105 = v81
	goto L27
L27:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v91+(v70+int32(4)))))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v70+v105<<(uint(int32(2))%32))))
	if v107 != v111 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v114 = v105 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v70+v114<<(uint(int32(2))%32)))) = v107
	v119 = v114
	goto L30
L29:
	;
	v119 = v105
	goto L30
L30:
	;
	v120 = int32(2)
	v121 = v83 + v120
	v123 = v84 + v120
	if v123 != v67&int32(-2) {
		v81 = v119
		v83 = v121
		v84 = v123
		goto L23
	} else {
		goto L31
	}
L31:
	;
	goto L24
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v152 + int32(1)
	goto L16
L33:
	;
	v137 = int32(2)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v70+v128<<(uint(v137)%32))))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v70+v126<<(uint(v137)%32))))
	if v140 == v144 {
		v152 = v126
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v147 = v126 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v70+v147<<(uint(int32(2))%32)))) = v140
	v152 = v147
	goto L32
}
func F_getPublicationSchemaInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = F_SearchSysCache1(m, int32(49), v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 == int32(0) {
			if l1 != 0 {
				v76 = int32(0)
				m.G0 = v10 + int32(32)
				return v76
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v25
					F_errmsg_internal(m, int32(55456), v10)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(493985), int32(2876), int32(242613))
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
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
			v37 = v35 + v36
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
			v39 = F_get_publication_name(m, v38, l1)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v39
				if v39 == int32(0) {
					F_ReleaseCatCache(m, v14)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						v76 = base.B2i32(v39 != int32(0))
						m.G0 = v10 + int32(32)
						return v76
					}
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
					v45 = F_get_namespace_name(m, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v45
						if v45 != 0 {
							F_ReleaseCatCache(m, v14)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								v76 = base.B2i32(v39 != int32(0))
								m.G0 = v10 + int32(32)
								return v76
							}
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							F_pfree(m, v49)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								F_ReleaseCatCache(m, v14)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									if l1 != 0 {
										v76 = int32(0)
										m.G0 = v10 + int32(32)
										return v76
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v48
											F_errmsg_internal(m, int32(55422), v10+int32(16))
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(493985), int32(2897), int32(242613))
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
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
				}
			}
		}
	}
}
func F_get_publication_name(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_SearchSysCache1(m, int32(51), l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			if l1 != 0 {
				v40 = int32(0)
				m.G0 = v8 + int32(16)
				return v40
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(46769), v8)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(499479), int32(3796), int32(379236))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
			v35 = F_pstrdup(m, v30+v31+int32(4))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v11)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v40 = v35
					m.G0 = v8 + int32(16)
					return v40
				}
			}
		}
	}
}
