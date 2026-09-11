package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_has_sequence_privilege_id_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v18)
		v21 = F_convert_any_priv_string(m, v14, int32(_a_F_has_sequence_privilege_id_id_0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = F_get_rel_relkind(m, v11)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 != int32(83) {
					if v23&int32(255) == int32(0) {
						v31 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
						v65 = int32(0)
						m.G0 = v9 + int32(16)
						return v65
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(151027844))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = F_get_rel_name(m, v11)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v41
									F_errmsg(m, int32(_a_F_has_sequence_privilege_id_id_1), v9)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_has_sequence_privilege_id_id_2), int32(2284), int32(_a_F_has_sequence_privilege_id_id_3))
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
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
				} else {
					v54 = F_pg_class_aclcheck_ext(m, v11, v12, v21, v9+int32(15))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
						if v56 == int32(1) {
							v59 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
							v65 = int32(0)
						} else {
							v65 = base.B2i32(v54 == int32(0))
						}
						m.G0 = v9 + int32(16)
						return v65
					}
				}
			}
		}
	}
}
func F_has_sequence_privilege_name_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v18)
		v20 = F_get_role_oid_or_public(m, v12)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v23 = F_convert_any_priv_string(m, v14, int32(_a_F_has_sequence_privilege_name_id_0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = F_get_rel_relkind(m, v11)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v25 != int32(83) {
						if v25&int32(255) == int32(0) {
							v33 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
							v67 = int32(0)
							m.G0 = v9 + int32(16)
							return v67
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(151027844))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v43 = F_get_rel_name(m, v11)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v43
										F_errmsg(m, int32(_a_F_has_sequence_privilege_name_id_1), v9)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_has_sequence_privilege_name_id_2), int32(2186), int32(_a_F_has_sequence_privilege_name_id_3))
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
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
					} else {
						v56 = F_pg_class_aclcheck_ext(m, v11, v20, v23, v9+int32(15))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
							if v58 == int32(1) {
								v61 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v61)
								v67 = int32(0)
							} else {
								v67 = base.B2i32(v56 == int32(0))
							}
							m.G0 = v9 + int32(16)
							return v67
						}
					}
				}
			}
		}
	}
}
func F_has_sequence_privilege_name_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v34 int32
	_ = v34
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
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = F_get_role_oid_or_public(m, v10)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v22 = F_convert_any_priv_string(m, v17, int32(_a_F_has_sequence_privilege_name_name_0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = F_textToQualifiedNameList(m, v12)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = F_makeRangeVarFromNameList(m, v24)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							v28 = int32(0)
							v32 = F_RangeVarGetRelidExtended(m, v26, v28, v28, v28, v28)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v34 = F_get_rel_relkind(m, v32)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									if v34 != int32(83) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v41 = m.ExcPending
										if v41 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(151027844))
											mBase = m.M
											v44 = m.ExcPending
											if v44 != 0 {
												return int32(0)
											} else {
												v45 = F_text_to_cstring(m, v12)
												mBase = m.M
												v46 = m.ExcPending
												if v46 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = v45
													F_errmsg(m, int32(_a_F_has_sequence_privilege_name_name_1), v8)
													mBase = m.M
													v50 = m.ExcPending
													if v50 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_has_sequence_privilege_name_name_2), int32(2123), int32(_a_F_has_sequence_privilege_name_name_3))
														mBase = m.M
														v55 = m.ExcPending
														if v55 != 0 {
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
									} else {
										v56 = F_pg_class_aclcheck(m, v32, v19, v22)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											m.G0 = v8 + int32(16)
											return base.B2i32(v56 == int32(0))
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
func F_init_sequence(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int64
	_ = v76
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l0
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_init_sequence[0]))
	if v13 == int32(0) {
		*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(171798691844)
		v24 = F_hash_create(m, int32(_a_F_init_sequence_0), int32(16), v7+int32(-48), int32(40))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_init_sequence[0])) = v24
			v27 = v24
			v33 = F_hash_search(m, v27, v7+int32(-52), int32(1), v7+int32(-48))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
				if v35 == int32(1) {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
					v48 = v38
				} else {
					v39 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v33)+4)) = v39
					*(*int64)(unsafe.Add(mBase, uint32(v33)+16)) = v39
					v43 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v33)+12)) = uint8(v43)
					*(*int64)(unsafe.Add(mBase, uint32(v33)+24)) = v39
					v48 = v43
				}
				v50 = *(*int32)(unsafe.Add(mBase, _c_F_init_sequence[1]))
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+56))
				if v48 != v51 {
					v53 = int32(_a_F_init_sequence_1)
					v54 = *(*int32)(unsafe.Add(mBase, _c_F_init_sequence[2]))
					v57 = *(*int32)(unsafe.Add(mBase, _c_F_init_sequence[3]))
					*(*int32)(unsafe.Add(mBase, _c_F_init_sequence[2])) = v57
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
					F_LockRelationOid(m, v59, int32(3))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_init_sequence[2])) = v54
						*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v51
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						v69 = F_sequence_open(m, v67, int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+88))
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
							if v72 != v73 {
								*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v72
								v76 = *(*int64)(unsafe.Add(mBase, uint32(v33)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v33)+24)) = v76
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v33
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v69
							m.G0 = v9 - int32(-64)
							return
						}
					}
				} else {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
					v69 = F_sequence_open(m, v67, int32(0))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+88))
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
						if v72 != v73 {
							*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v72
							v76 = *(*int64)(unsafe.Add(mBase, uint32(v33)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v33)+24)) = v76
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v33
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v69
						m.G0 = v9 - int32(-64)
						return
					}
				}
			}
		}
	} else {
		v27 = v13
		v33 = F_hash_search(m, v27, v7+int32(-52), int32(1), v7+int32(-48))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
			if v35 == int32(1) {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
				v48 = v38
			} else {
				v39 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v33)+4)) = v39
				*(*int64)(unsafe.Add(mBase, uint32(v33)+16)) = v39
				v43 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v33)+12)) = uint8(v43)
				*(*int64)(unsafe.Add(mBase, uint32(v33)+24)) = v39
				v48 = v43
			}
			v50 = *(*int32)(unsafe.Add(mBase, _c_F_init_sequence[1]))
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+56))
			if v48 != v51 {
				v53 = int32(_a_F_init_sequence_1)
				v54 = *(*int32)(unsafe.Add(mBase, _c_F_init_sequence[2]))
				v57 = *(*int32)(unsafe.Add(mBase, _c_F_init_sequence[3]))
				*(*int32)(unsafe.Add(mBase, _c_F_init_sequence[2])) = v57
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
				F_LockRelationOid(m, v59, int32(3))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_init_sequence[2])) = v54
					*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v51
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
					v69 = F_sequence_open(m, v67, int32(0))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+88))
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
						if v72 != v73 {
							*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v72
							v76 = *(*int64)(unsafe.Add(mBase, uint32(v33)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v33)+24)) = v76
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v33
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v69
						m.G0 = v9 - int32(-64)
						return
					}
				}
			} else {
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
				v69 = F_sequence_open(m, v67, int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+88))
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
					if v72 != v73 {
						*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v72
						v76 = *(*int64)(unsafe.Add(mBase, uint32(v33)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v33)+24)) = v76
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v33
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v69
					m.G0 = v9 - int32(-64)
					return
				}
			}
		}
	}
}
func F_sequence_open(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_relation_open(m, l0, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+119)))
		if v13 != int32(83) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v23 + int32(4)
					F_errmsg(m, int32(_a_F_sequence_open_0), v6)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
						v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30)+119)))
						F_errdetail_relkind_not_supported(m, v31)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_sequence_open_1), int32(77), int32(_a_F_sequence_open_2))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
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
		} else {
			m.G0 = v6 + int32(16)
			return v8
		}
	}
}
