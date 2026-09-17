package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_performDeletion(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v13
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v17 - int32(1259) {
		case 0:
			F_LockRelationOid(m, v16, int32(8))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v31 = F_palloc(m, int32(16))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = int64(137438953472)
					v36 = F_palloc(m, int32(384))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v38
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = v36
						v45 = v9 + int32(12)
						F_findDependentObjects(m, l0, int32(1), l2, v38, v31, v38, v45)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							F_reportDependentObjects(m, v31, l1, l2, l0)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								F_deleteObjectsInList(m, v31, v45, l2)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
									F_pfree(m, v52)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
										if v55 != 0 {
											F_pfree(m, v55)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return
											} else {
												F_pfree(m, v31)
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return
												} else {
													v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
													F_relation_close(m, v60, int32(3))
													mBase = m.M
													v63 = m.ExcPending
													if v63 != 0 {
														return
													} else {
														m.G0 = v9 + int32(16)
														return
													}
												}
											}
										} else {
											F_pfree(m, v31)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return
											} else {
												v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												F_relation_close(m, v60, int32(3))
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return
												} else {
													m.G0 = v9 + int32(16)
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
		default:
			F_LockDatabaseObject(m, v17, v16, int32(8))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v31 = F_palloc(m, int32(16))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = int64(137438953472)
					v36 = F_palloc(m, int32(384))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v38
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = v36
						v45 = v9 + int32(12)
						F_findDependentObjects(m, l0, int32(1), l2, v38, v31, v38, v45)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							F_reportDependentObjects(m, v31, l1, l2, l0)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								F_deleteObjectsInList(m, v31, v45, l2)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
									F_pfree(m, v52)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
										if v55 != 0 {
											F_pfree(m, v55)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return
											} else {
												F_pfree(m, v31)
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return
												} else {
													v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
													F_relation_close(m, v60, int32(3))
													mBase = m.M
													v63 = m.ExcPending
													if v63 != 0 {
														return
													} else {
														m.G0 = v9 + int32(16)
														return
													}
												}
											}
										} else {
											F_pfree(m, v31)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return
											} else {
												v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												F_relation_close(m, v60, int32(3))
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return
												} else {
													m.G0 = v9 + int32(16)
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
		case 2:
			F_LockSharedObject(m, int32(1261), v16, int32(8))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v31 = F_palloc(m, int32(16))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = int64(137438953472)
					v36 = F_palloc(m, int32(384))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v38
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = v36
						v45 = v9 + int32(12)
						F_findDependentObjects(m, l0, int32(1), l2, v38, v31, v38, v45)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							F_reportDependentObjects(m, v31, l1, l2, l0)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								F_deleteObjectsInList(m, v31, v45, l2)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
									F_pfree(m, v52)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
										if v55 != 0 {
											F_pfree(m, v55)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return
											} else {
												F_pfree(m, v31)
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return
												} else {
													v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
													F_relation_close(m, v60, int32(3))
													mBase = m.M
													v63 = m.ExcPending
													if v63 != 0 {
														return
													} else {
														m.G0 = v9 + int32(16)
														return
													}
												}
											}
										} else {
											F_pfree(m, v31)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return
											} else {
												v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												F_relation_close(m, v60, int32(3))
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return
												} else {
													m.G0 = v9 + int32(16)
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
}
