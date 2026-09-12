package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AppendStringToManifest(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = F_strlen(m, l1)
	mBase = m.M
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v10 != int32(1) {
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_BufFileWrite(m, v46, l1, v9)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v49 + base.I64_extend_i32_s(v9)
			m.G0 = v7 + int32(16)
			return
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = F_pg_cryptohash_update(m, v13, l1, v9)
		mBase = m.M
		if int32(0) <= v14 {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			F_BufFileWrite(m, v46, l1, v9)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v49 + base.I64_extend_i32_s(v9)
				m.G0 = v7 + int32(16)
				return
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v21 == int32(0) {
					v36 = int32(14012)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
					if v28 == int32(1) {
						v31 = int32(316782)
					} else {
						v31 = int32(137542)
					}
					if v28 == int32(2) {
						v34 = int32(14012)
					} else {
						v34 = v31
					}
					v36 = v34
				}
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v36
				F_errmsg_internal(m, int32(208345), v7)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_errfinish(m, int32(511925), int32(393), int32(83249))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
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
func F_appendStringInfoChar(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = l1
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v4 <= v5+int32(1) {
		F_enlargeStringInfo(m, l0, int32(1))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v13 = v12
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*uint8)(unsafe.Add(mBase, uint32(v13+v14))) = uint8(v2)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v19 = v17 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v23 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v21+v19))) = uint8(v23)
			return
		}
	} else {
		v13 = v5
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*uint8)(unsafe.Add(mBase, uint32(v13+v14))) = uint8(v2)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v19 = v17 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v23 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v21+v19))) = uint8(v23)
		return
	}
}
func F_convert_string_datum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	v1 = l0
	switch l1 - int32(18) {
	case 0:
		v20 = F_palloc(m, int32(2))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v1)
			v23 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)) = uint8(v23)
			v27 = v20
			v28 = F_pg_newlocale_from_collation(m, l2)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
				if v30 != 0 {
					return v27
				} else {
					v32 = int32(0)
					v34 = F_pg_strxfrm(m, v32, v27, v32, v28)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v37 = v34 + int32(1)
						v38 = F_palloc(m, v37)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = F_pg_strxfrm(m, v38, v27, v37, v28)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v27)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									return v38
								}
							}
						}
					}
				}
			}
		}
	case 1:
		v7 = F_pstrdup(m, v1)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v27 = v7
			v28 = F_pg_newlocale_from_collation(m, l2)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
				if v30 != 0 {
					return v27
				} else {
					v32 = int32(0)
					v34 = F_pg_strxfrm(m, v32, v27, v32, v28)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v37 = v34 + int32(1)
						v38 = F_palloc(m, v37)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = F_pg_strxfrm(m, v38, v27, v37, v28)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v27)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									return v38
								}
							}
						}
					}
				}
			}
		}
	case 2, 3, 4, 5, 6:
		v15 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v15)
		return int32(0)
	case 7:
		v25 = F_text_to_cstring(m, v1)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v27 = v25
			v28 = F_pg_newlocale_from_collation(m, l2)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
				if v30 != 0 {
					return v27
				} else {
					v32 = int32(0)
					v34 = F_pg_strxfrm(m, v32, v27, v32, v28)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v37 = v34 + int32(1)
						v38 = F_palloc(m, v37)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = F_pg_strxfrm(m, v38, v27, v37, v28)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v27)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									return v38
								}
							}
						}
					}
				}
			}
		}
	default:
		if base.Ui32(l1-int32(1042)) < base.Ui32(int32(2)) {
			v25 = F_text_to_cstring(m, v1)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = v25
				v28 = F_pg_newlocale_from_collation(m, l2)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
					if v30 != 0 {
						return v27
					} else {
						v32 = int32(0)
						v34 = F_pg_strxfrm(m, v32, v27, v32, v28)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v37 = v34 + int32(1)
							v38 = F_palloc(m, v37)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = F_pg_strxfrm(m, v38, v27, v37, v28)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v27)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return int32(0)
									} else {
										return v38
									}
								}
							}
						}
					}
				}
			}
		} else {
			v15 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v15)
			return int32(0)
		}
	}
}
func F_makeString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(8))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(468)
		return v4
	}
}
func F_makeStringConstCast(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v4 = F_makeStringConst(m, l0, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = F_palloc0(m, int32(16))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v4
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(73)
			return v9
		}
	}
}
func F_string_agg_combine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
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
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = v8 + int32(8)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 == v2 {
		v30 = int32(0)
		if v11 == v30 {
			v38 = v30
		} else {
			v33 = v30
			v34 = v2
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v33
			v38 = v34
		}
		v41 = v38
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		switch v16 - int32(429) {
		case 0:
			if v11 == int32(0) {
				v41 = int32(1)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+168))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
				v33 = v23
				v34 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v33
				v38 = v34
				v41 = v38
			}
		case 1:
			if v11 == int32(0) {
				v41 = int32(2)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+368))
				v33 = v28
				v34 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v33
				v38 = v34
				v41 = v38
			}
		default:
			v30 = int32(0)
			if v11 == v30 {
				v38 = v30
			} else {
				v33 = v30
				v34 = v2
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v33
				v38 = v34
			}
			v41 = v38
		}
	}
	if v41 != 0 {
		v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v42 == int32(0) {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v46 = v45
		} else {
			v46 = v2
		}
		v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v47 == int32(0) {
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v50 != 0 {
				if v46 == int32(0) {
					v57 = int32(4549024)
					v58 = *(*int32)(unsafe.Add(mBase, _consts[10]))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
					*(*int32)(unsafe.Add(mBase, _consts[10])) = v60
					v63 = v8 + int32(12)
					v64 = int32(0)
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v65 == v64 {
						v82 = int32(0)
						if v63 == v82 {
							v90 = v82
						} else {
							v85 = v82
							v86 = v64
							*(*int32)(unsafe.Add(mBase, uint32(v63))) = v85
							v90 = v86
						}
						v93 = v90
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
						switch v68 - int32(429) {
						case 0:
							if v63 == int32(0) {
								v93 = int32(1)
							} else {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+168))
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
								v85 = v75
								v86 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v63))) = v85
								v90 = v86
								v93 = v90
							}
						case 1:
							if v63 == int32(0) {
								v93 = int32(2)
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v65)+368))
								v85 = v80
								v86 = int32(2)
								*(*int32)(unsafe.Add(mBase, uint32(v63))) = v85
								v90 = v86
								v93 = v90
							}
						default:
							v82 = int32(0)
							if v63 == v82 {
								v90 = v82
							} else {
								v85 = v82
								v86 = v64
								*(*int32)(unsafe.Add(mBase, uint32(v63))) = v85
								v90 = v86
							}
							v93 = v90
						}
					}
					if v93 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v145 = m.ExcPending
						if v145 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(66755), int32(0))
							mBase = m.M
							v149 = m.ExcPending
							if v149 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(521251), int32(5432), int32(368767))
								mBase = m.M
								v154 = m.ExcPending
								if v154 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v96 = int32(4549024)
						v97 = *(*int32)(unsafe.Add(mBase, _consts[10]))
						v99 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						*(*int32)(unsafe.Add(mBase, _consts[10])) = v99
						v101 = F_makeStringInfo(m)
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[10])) = v97
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
							F_appendBinaryStringInfo(m, v101, v107, v108)
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return int32(0)
							} else {
								v111 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v101)+12)) = v111
								*(*int32)(unsafe.Add(mBase, _consts[10])) = v58
								v122 = v101
								m.G0 = v8 + int32(16)
								return v122
							}
						}
					}
				} else {
					v115 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
					if v115 <= int32(0) {
						v122 = v46
						m.G0 = v8 + int32(16)
						return v122
					} else {
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
						F_appendBinaryStringInfo(m, v46, v118, v115)
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							v122 = v46
							m.G0 = v8 + int32(16)
							return v122
						}
					}
				}
			} else {
				if v46 != 0 {
					v122 = v46
				} else {
					v52 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v52)
					v122 = int32(0)
				}
				m.G0 = v8 + int32(16)
				return v122
			}
		} else {
			if v46 != 0 {
				v122 = v46
			} else {
				v52 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v52)
				v122 = int32(0)
			}
			m.G0 = v8 + int32(16)
			return v122
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v132 = m.ExcPending
		if v132 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(66371), int32(0))
			mBase = m.M
			v136 = m.ExcPending
			if v136 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(521251), int32(5509), int32(389244))
				mBase = m.M
				v141 = m.ExcPending
				if v141 != 0 {
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
func F_string_compare(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	v5 = l2 - int32(1)
	if v5 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v49
L2:
	;
	v49 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v12 = l0
	v13 = l1
	v14 = v5
	v15 = v11
	goto L9
L6:
	;
	v37 = l1
	v41 = int32(0)
	goto L7
L7:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	v49 = v41 - v42
	goto L1
L8:
	;
	v37 = v32
	v41 = v34
	goto L7
L9:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v15 != v17 {
		v32 = v13
		v34 = v15
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v32 = v26
	v34 = int32(0)
	goto L8
L11:
	;
	if v17 == int32(0) {
		v32 = v13
		v34 = v15
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v22 = v14 - int32(1)
	if v22 == int32(0) {
		v32 = v13
		v34 = v15
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v25 = int32(1)
	v26 = v13 + v25
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v27 != 0 {
		v12 = v12 + v25
		v13 = v26
		v14 = v22
		v15 = v27
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
}
