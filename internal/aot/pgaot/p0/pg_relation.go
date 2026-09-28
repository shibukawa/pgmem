package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ScanPgRelation(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v34 int32
	_ = v34
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ScanPgRelation[0]))
	if v12 != 0 {
		F_ScanKeyInit(m, v9, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v23 = F_table_open(m, int32(1259), int32(1))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if l2 != 0 {
					v26 = F_GetNonHistoricCatalogSnapshot(m, int32(1259))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = F_RegisterSnapshot(m, v26)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v30 = v28
							v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ScanPgRelation[1])))
							v37 = F_systable_beginscan(m, v23, int32(2662), l1&v34, v30, int32(1), v9)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v39 = F_systable_getnext(m, v37)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									if v39 != 0 {
										v41 = F_heap_copytuple(m, v39)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return int32(0)
										} else {
											v43 = v41
											F_systable_endscan(m, v37)
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return int32(0)
											} else {
												if v30 != 0 {
													F_UnregisterSnapshot(m, v30)
													mBase = m.M
													v47 = m.ExcPending
													if v47 != 0 {
														return int32(0)
													} else {
														F_relation_close(m, v23, int32(1))
														mBase = m.M
														v50 = m.ExcPending
														if v50 != 0 {
															return int32(0)
														} else {
															m.G0 = v9 - int32(-64)
															return v43
														}
													}
												} else {
													F_relation_close(m, v23, int32(1))
													mBase = m.M
													v50 = m.ExcPending
													if v50 != 0 {
														return int32(0)
													} else {
														m.G0 = v9 - int32(-64)
														return v43
													}
												}
											}
										}
									} else {
										v43 = int32(0)
										F_systable_endscan(m, v37)
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return int32(0)
										} else {
											if v30 != 0 {
												F_UnregisterSnapshot(m, v30)
												mBase = m.M
												v47 = m.ExcPending
												if v47 != 0 {
													return int32(0)
												} else {
													F_relation_close(m, v23, int32(1))
													mBase = m.M
													v50 = m.ExcPending
													if v50 != 0 {
														return int32(0)
													} else {
														m.G0 = v9 - int32(-64)
														return v43
													}
												}
											} else {
												F_relation_close(m, v23, int32(1))
												mBase = m.M
												v50 = m.ExcPending
												if v50 != 0 {
													return int32(0)
												} else {
													m.G0 = v9 - int32(-64)
													return v43
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v30 = int32(0)
					v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ScanPgRelation[1])))
					v37 = F_systable_beginscan(m, v23, int32(2662), l1&v34, v30, int32(1), v9)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = F_systable_getnext(m, v37)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v39 != 0 {
								v41 = F_heap_copytuple(m, v39)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v43 = v41
									F_systable_endscan(m, v37)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										if v30 != 0 {
											F_UnregisterSnapshot(m, v30)
											mBase = m.M
											v47 = m.ExcPending
											if v47 != 0 {
												return int32(0)
											} else {
												F_relation_close(m, v23, int32(1))
												mBase = m.M
												v50 = m.ExcPending
												if v50 != 0 {
													return int32(0)
												} else {
													m.G0 = v9 - int32(-64)
													return v43
												}
											}
										} else {
											F_relation_close(m, v23, int32(1))
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return int32(0)
											} else {
												m.G0 = v9 - int32(-64)
												return v43
											}
										}
									}
								}
							} else {
								v43 = int32(0)
								F_systable_endscan(m, v37)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									if v30 != 0 {
										F_UnregisterSnapshot(m, v30)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return int32(0)
										} else {
											F_relation_close(m, v23, int32(1))
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return int32(0)
											} else {
												m.G0 = v9 - int32(-64)
												return v43
											}
										}
									} else {
										F_relation_close(m, v23, int32(1))
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return int32(0)
										} else {
											m.G0 = v9 - int32(-64)
											return v43
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
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_ScanPgRelation_0), int32(0))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_ScanPgRelation_1), int32(359), int32(_a_F_ScanPgRelation_2))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
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
func F_pg_relation_is_updatable(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int32(0)
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = F_relation_is_updatable(m, v2, v3, base.B2i32(v4 != int64(0)), v3)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_s(v8)
	}
}
